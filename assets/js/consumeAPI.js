const urlApi = "http://localhost:8888/sertificate/all";
const tbody = document.getElementById("tbody");
let data = "";
const GetData = (posts) => {
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

const postData = () => {
    return{
        name: document.getElementById("nama").value,
        kecamatan: document.getElementById("kecamatan").value,
        pelatihan: document.getElementById("pelatihan").value,
        keterangan: document.getElementById("keterangan").value
    }
};

// get data
fetch(urlApi)
.then(response => response.json())
.then(data => {
    const realData = data.data;
    console.log(realData);
    GetData(realData);
});

// post data (create data)
const addForm = document.querySelector('.popup-body');
addForm.addEventListener('submit', (e) => {
    e.preventDefault();
    let payload = postData();
    fetch("http://localhost:8888/sertificate", {
        method: 'POST',
        headers:{
            "Content-Type":"application/json"
        },
        body: JSON.stringify(payload)
    })
    .then(response => response.json())
    .then(data => {
        const dataArr = [];
        dataArr.push(data.data);
        GetData(dataArr);
        location.reload();
    })
});