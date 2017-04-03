var clipboard = new Clipboard('.copyButton');
clipboard.on('success', function(e) {
    showInfoMessage("Copied!");
});

function showSuccessMessage(message){
    $.notify(message, "success");
}

function showInfoMessage(message){
    $.notify(message, "info");
}

function showWarningMessage(message){
    $.notify(message,"warn");
}

function showErrorMessage(message) {
    $.notify(message, "error");
}