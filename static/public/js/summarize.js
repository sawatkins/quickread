window.onload = function () {
    document.getElementById("form-upload-doc").addEventListener("submit", handleDocUploadSummary);
    document.getElementById("input-upload-doc").addEventListener('change', function () {
        if (this.value && isProperFile()) {
            document.getElementById('button-upload-doc').disabled = false;
            displayFileName(true);
        } else {
            document.getElementById('button-upload-doc').disabled = true;
            displayFileName(false);
        }
    });
}

function handleDocUploadSummary(event) {
    displaySummaryResultPage();
    event.preventDefault();
    const formData = new FormData(event.target);
    console.log(formData)
    uploadDoc(formData).then(data => {
        if (data.presignedUrl) {
            summarizeDoc(data.presignedUrl);
        } else {
            hideSpinner();
            document.getElementById("summary-response-text").innerText = "File upload filed. Please try again."
        }
    }).catch(error => {
        console.error(error)
    })
}

async function uploadDoc(formData) {
    const response = await fetch("/upload-doc", {
        method: "post",
        body: formData
    });
    // add check if response is ok and only return data is success
    const data = await response.json();
    return data;
}

async function summarizeDoc(presignedUrl) {
    const summaryResonseElement = document.getElementById("summary-response-text");
    let params = new URLSearchParams();
    params.append("presignedUrl", presignedUrl)
    console.log("/summarize-doc?" + params)
    fetch("/summarize-doc?" + params)
        .then(response => response.json())
        .then(data => {
            hideSpinner();
            summaryResonseElement.innerText = data;
        })
        .catch(error => {
            console.log(error)
        })
}

function isProperFile() {
    const fileInput = document.getElementById("input-upload-doc");
    const fileName = fileInput.files[0].name;
    const fileSize = fileInput.files[0].size;
    const fileSizeInMB = Math.round(fileSize / (1024 * 1024));
    if (fileSizeInMB > 50) {
        alert('File size must be less than or equal to 50MB');
        return false;
    }
    if (!fileName.endsWith('.pdf')) {
        alert('Please select a PDF file');
        return false;
    }
    return true;
}

function displaySummaryResultPage() {
    hideAllSummaryBlocks();
    const fileInput = document.getElementById('input-upload-doc');
    const fileName = document.getElementById('file-summary-name-response');
    fileName.innerText = fileInput.files[0].name;
    showElementByID("summary-response");
    displaySpinner();
}

function displaySpinner() {
    const spinner = document.getElementById('summary-spinner');
    spinner.style.display = 'block';
}

function hideSpinner() {
    const spinner = document.getElementById('summary-spinner');
    spinner.style.display = 'none';
}

function goToSummaryUploadPage() {
    hideAllSummaryBlocks();
    document.getElementById('input-upload-doc').value = '';
    document.getElementById('upload-doc-filename').textContent = '';
    document.getElementById('summary-response').innerText = '';
    showElementByID("document");
}

function hideAllSummaryBlocks() {
    summaryBlocks = document.getElementsByClassName("summary-block");
    for (let i = 0; i < summaryBlocks.length; i++) {
        let block = summaryBlocks[i];
        block.style.display = 'none';
    }
}

function displayFileName(toDisplay) {
    const fileInput = document.getElementById('input-upload-doc');
    const fileName = document.getElementById('upload-doc-filename');
    if (toDisplay) {
        fileName.textContent = "File: " + fileInput.files[0].name;

    } else {
        fileName.textContent = '';
    }
}
