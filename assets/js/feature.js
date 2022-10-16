document.getElementById('button-addData').addEventListener('click', function() {
    document.querySelector('.popup').style.display = 'flex';
});

document.getElementById('close-button').addEventListener('click', function() {
    document.querySelector('.popup').style.display = 'none';
});

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