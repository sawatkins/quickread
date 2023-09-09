window.onload = function () {
    console.log("loaded summarize.js");

    // Submits form to /upload endpoint
    // document.getElementById("form-upload-doc").addEventListener("submit", async function (e) {
    // });

    const docUploadForm = document.getElementById("form-upload-doc");
    docUploadForm.addEventListener("submit", handleDocSummary);

}

async function handleDocUploadSummary(event) {
    displaySummaryResultPage();
    event.preventDefault();
    const formData = new FormData(event.target);

    console.log("formdata:", formData)

    //upload doc
    uploadDoc(formData).then(data => {
        console.log(data);
    }).catch(error => {
        console.error(error)
    })


}
// success = presigned url
// failure = signal ui to try again (maybe with errors; show upload page again...)
async function uploadDoc(formData) {
    const response = await fetch("/upload-doc", {
        method: "post",
        body: formData
    });
    // add check if response is ok and only return data is success
    const data = await response.json();
    return data;
}

async function displayPDFSummaryResponse() {
    console.log("instide displayPDFSummaryResponse()")
    const fileInput = document.getElementById('input-upload-doc');
    const params = new URLSearchParams();
    if (fileInput.files.length > 0) {
        params.append("filename", fileInput.files[0].name)
        fetch("/summarize_pdf?" + params, {
            method: "GET"
        })
            .then(response => response.json())
            .then(data => {
                console.log(data);
                const summaryResonseElement = document.getElementById("p-summary-response");
                summaryResonseElement.textContent = data;
            })
            .catch(error => {
                console.log(error)
            })
    }
}

function displaySummaryResultPage() {
    hideElementByID("document");
    hideElementByID("article");
    showElementByID("summary-response");
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
    const element = document.getElementById(element);
    element.style.display = "none";
}

function showElementByID(element) {
    const element = document.getElementById(element);
    element.style.display = "block";
}