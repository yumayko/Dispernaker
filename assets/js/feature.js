const openAddData = document.querySelectorAll(".button");
for (let i = 0; i < openAddData.length; i++) {
    openAddData[i].addEventListener("click", function () {
        document.querySelector('.popup').style.display = 'flex';

        if(document.querySelector('.popup').style.display === 'flex') {
            document.querySelector('body').style.overflow = 'hidden';
        }
    });
}

document.getElementById('close-button').addEventListener('click', function() {
    document.querySelector('.popup').style.display = 'none';
    
    if(document.querySelector('.popup').style.display === 'none') {
        document.querySelector('body').style.overflow = 'auto';
    }
});

const accordionItemHeaders = document.querySelectorAll(".acc-header");
console.log(accordionItemHeaders);
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