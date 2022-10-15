document.getElementById('button-addData').addEventListener('click', function() {
    document.querySelector('.popup').style.display = 'flex';
});

document.getElementById('close-button').addEventListener('click', function() {
    document.querySelector('.popup').style.display = 'none';
});

const according = document.getElementsByClassName('according-container');
for (let i = 0; i < according.length; i++) {
    according[i].addEventListener('click', function() {
        this.classList.toggle('active');
    });
}