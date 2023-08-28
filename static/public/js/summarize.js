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
                // toggleDisplayByElementID("div-doc-pages-selection");
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
    const summaryResponse = document.getElementById("div-summary-response");
    fetch("/summarize_pdf?" + params, {
        method: "GET"
    })
        .then(response => response.json())
        .then(data => {
            console.log(data);
        })
        .catch(error => {
            console.log(error)
        })

    // update the elemeent witht the text
}

async function getPresignedUrlFromS3() {
    const fileInput = document.getElementById('input-upload-doc');
    let presignUrl = "";
    if (fileInput.files.length > 0) {
        const params = new URLSearchParams();
        params.append("filename", fileInput.files[0].name);
        fetch('/get_presigned_url?' + params)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                return "";
            })
            .catch(error => {
                // Handle any errors
                console.error(error);
            });

    }
    return presignUrl;
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
    if (element.style.display === "none") {
        element.style.display = "block"; // maybe just have a way to remove the attribute instead
    } else {
        element.style.display = "none";
    }
}