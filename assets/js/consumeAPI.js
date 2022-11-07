const urlApi = "http://localhost:8888/sertificate/all";
const tbody = document.getElementById("tbody");
let data = "";
const postData = (posts) => {
    let i = 1;
    posts.forEach((post) => {
        data += 
        `<tr>
            <td>${i++}</td>
            <td>${post.name}</td>
            <td>${post.kecamatan}</td>
            <td>${post.pelatihan}</td>
            <td>${post.keterangan}</td>
            <td><button class="button-edit" id="button-edit">Edit</button></td>
            <td><button class="button-delete" id="button-delete">Hapus</button></td>
        </tr>`;
    })
    tbody.innerHTML = data;
};

// get data
fetch(urlApi)
.then(response => response.json())
.then(data => {
    const realData = data.data;
    console.log(realData);
    postData(realData);
});

// post data (create data)
const addForm = document.querySelector('.popup-body');
addForm.addEventListener('submit', (e) => {
    e.preventDefault();
    console.log('submit');
});