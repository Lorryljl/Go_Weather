function showProgress() {
    var progressBar = document.querySelector('.progress-bar');
    var progress = document.querySelector('.progress');
    progress.style.width = '100%';
    setTimeout(function() {
        progress.style.width = '0%';
    }, 2000);
}