window.onload = function () {
    console.log("loaded summarize.js");

    // Submits form to /upload endpoint
    // document.getElementById("form-upload-doc").addEventListener("submit", async function (e) {
    // });

    const docUploadForm = document.getElementById("form-upload-doc");
    docUploadForm.addEventListener("submit", handleDocUploadSummary);

}

function handleDocUploadSummary(event) {
    displaySummaryResultPage();
    event.preventDefault();
    const formData = new FormData(event.target);
    // console.log(formData)
    uploadDoc(formData).then(data => {
        console.log(data);
        // TODO if presigned url is undefuned, do something else
        summarizeDoc(data.presignedUrl)
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
    console.log("inside sumamrize doc")
    const summaryResonseElement = document.getElementById("summary-response");
    let params = new URLSearchParams();
    params.append("presignedUrl", presignedUrl)
    console.log("/summarize-doc?" + params)
    fetch("/summarize-doc?" + params)
        .then(response => response.json())
        .then(data => {
            summaryResonseElement.innerText = data;
        })
        .catch(error => {
            console.log(error)
        })
}

function displaySummaryLoadingPage() {
    
}

function displaySummaryResultPage() {
    // const fileInput = document.getElementById('input-upload-doc');
    // const fileName = document.getElementById('file-summary-name');

    hideElementByID("document");
    hideElementByID("article");
    // fileName.innerText = fileInput.files[0].name;
    showElementByID("summary-response");
}

function hideSummaryResultPage() {
    showElementByID("document");
    showElementByID("article");
    hideElementByID("summary-response");
}

// Displays succes message on upload success
function displayUploadSuccess() {
    // const resultStatusText = document.getElementById("upload-result-status");
    // resultStatusText.textContent = "File uploaded successfully"
}

// Displays failure and error message on upload failure
function displayUploadFailure(error) {
    // const resultStatusText = document.getElementById("upload-result-status");
    // const resultStatusTextError = document.getElementById("upload-result-status-error");
    // resultStatusText.textContent = "File upload failed. Please try again or submit a bug report.";
    // resultStatusTextError.textContent = error;
}

// Clears the upload result status text
function clearUploadResult() {
    // const resultStatusText = document.getElementById("upload-result-status");
    // const resultStatusTextError = document.getElementById("upload-result-status-error")
    // resultStatusText.textContent = "";
    // resultStatusTextError.textContent = "";
}

// Displays the file name in the UI after user selection
function displayFileName() {
    const fileInput = document.getElementById('input-upload-doc');
    const fileName = document.getElementById('upload-doc-filename');

    if (fileInput.files.length > 0) {
        fileName.textContent = "File: " + fileInput.files[0].name;
    } else {
        fileName.textContent = '';
    }
}

function toggleDisplayByID(elementToToggle) {
    const element = document.getElementById(elementToToggle);
    if (element.style.display == "none") {
        element.style.display = "block"; // maybe just have a way to remove the attribute instead
    } else {
        element.style.display = "none";
    }
}

function hideElementByID(element) {
    const elementToHide = document.getElementById(element);
    elementToHide.style.display = "none";
}

function showElementByID(element) {
    const elementToShow = document.getElementById(element);
    elementToShow.style.display = "block";
}