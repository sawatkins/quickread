window.onload = function () {

    document.getElementById("form-upload-doc").addEventListener("submit", function (e) {
        e.preventDefault();

        // Handle form submission here, e.g., upload the file to a server
        console.log("Form submitted");
    });

    console.log("loaded main.js");

}

function displayFileName() {
    const fileInput = document.getElementById('input-upload-doc');
    const fileName = document.getElementById('upload-doc-filename');

    if (fileInput.files.length > 0) {
        fileName.textContent = "File: " + fileInput.files[0].name;
    } else {
        fileName.textContent = '';
    }
}