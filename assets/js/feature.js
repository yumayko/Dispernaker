const openAddData = document.querySelectorAll(".button");
for (let i = 0; i < openAddData.length; i++) {
    openAddData[i].addEventListener("click", function () {
        document.querySelector('.popup').style.display = 'flex';

        if(document.querySelector('.popup').style.display === 'flex') {
            document.querySelector('body').style.overflow = 'hidden';
        }
    });
}

document.querySelector(".calonPel").onclick = function() {
    document.getElementById("submit-pesertaPel").style.display = "none";
    document.getElementById("submit-calonSert").style.display = "none";
    document.getElementById("submit-pesertaSert").style.display = "none";
}
document.querySelector(".pesertaPel").onclick = function() {
    document.getElementById("submit-calonPel").style.display = "none";
    document.getElementById("submit-calonSert").style.display = "none";
    document.getElementById("submit-pesertaSert").style.display = "none";
}
document.querySelector(".calonSert").onclick = function() {
    document.getElementById("submit-calonPel").style.display = "none";
    document.getElementById("submit-pesertaPel").style.display = "none";
    document.getElementById("submit-pesertaSert").style.display = "none";
}
document.querySelector(".pesertaSert").onclick = function() {
    document.getElementById("submit-calonPel").style.display = "none";
    document.getElementById("submit-pesertaPel").style.display = "none";
    document.getElementById("submit-calonSert").style.display = "none";
}

const closeAddData = document.querySelectorAll(".close-button");
for (let i = 0; i < closeAddData.length; i++) {
    closeAddData[i].addEventListener("click", function () {
        const popup = document.querySelector('.popup').style.display = 'none';
        const popupEdit = document.querySelector('.popupEdit').style.display = 'none';
        
        if(popup === 'none' || popupEdit === 'none') {
            document.querySelector('body').style.overflow = 'auto';
            document.getElementById("submit-calonPel").style.display = "block";
            document.getElementById("submit-pesertaPel").style.display = "block";
            document.getElementById("submit-calonSert").style.display = "block";
            document.getElementById("submit-pesertaSert").style.display = "block";
        }
    });
}

const accordionItemHeaders = document.querySelectorAll(".acc-header");
accordionItemHeaders.forEach(accordionItemHeader => {
    accordionItemHeader.addEventListener("click", event => {
    accordionItemHeader.classList.toggle("active");
    const accordionItemBody = accordionItemHeader.nextElementSibling;
    if(accordionItemHeader.classList.contains("active")) {
        accordionItemBody.style.maxHeight = accordionItemBody.scrollHeight + "px";
    }
    else {
        accordionItemBody.style.maxHeight = 0;
    }
    
    });
});