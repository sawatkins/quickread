window.onload = function () {
    console.log("loaded upload.js");
    document.getElementById("form-upload-doc").addEventListener("submit", async function (e) {
        e.preventDefault();

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


function displayFileName() {
    const fileInput = document.getElementById('input-upload-doc');
    const fileName = document.getElementById('upload-doc-filename');

    if (fileInput.files.length > 0) {
        fileName.textContent = "File: " + fileInput.files[0].name;
    } else {
        fileName.textContent = '';
    }
}