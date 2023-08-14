window.onload = function() {
    
    document.getElementById("form-upload-doc").addEventListener("submit", function (e) {
        e.preventDefault();
    
        // Handle form submission here, e.g., upload the file to a server
        console.log("Form submitted");
    });
    
    console.log("loaded main.js");

}