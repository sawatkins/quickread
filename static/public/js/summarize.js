window.onload = function () {
    console.log("loaded summarize.js");

    clearUploadResult();

    // Submits form to /upload endpoint
    document.getElementById("form-upload-doc").addEventListener("submit", async function (e) {
        e.preventDefault();
        clearUploadResult();

        const form = e.target;
        const url = form.action;
        const formData = new FormData(form);

        try {
            const response = await fetch(url, {
                method: form.method,
                body: formData
            });

            if (response.ok) {
                const result = await response.json();
                // toggleDisplayByElementID("div-summary-response");
                console.log('File uploaded successfully:', result);
            } else {
                console.error('File upload failed:', response.statusText);
            }
        } catch (error) {
            console.error('Error uploading file:', error);
        }

        console.log("Form submitted");
    });

}

//
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



function toggleDisplayByElementID(elementToToggle) {
    const element = document.getElementById(elementToToggle);
    if (element.style.display == "none") {
        element.style.display = "block"; // maybe just have a way to remove the attribute instead
    } else {
        element.style.display = "none";
    }
}