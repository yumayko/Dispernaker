const tbody = document.getElementById("tbody");
const Vname = document.getElementById("nama");
const Vkecamatan = document.getElementById("kecamatan");
const Vpelatihan = document.getElementById("pelatihan");
const Vketerangan = document.getElementById("keterangan");
let data = "";
const GetData = (posts) => {
    let i = 1;
    posts.forEach((post) => {
        data += 
        `<tr data-id=${post.id}>
            <td>${i++}</td>
            <td class="name">${post.name}</td>
            <td class="kecamatan">${post.kecamatan}</td>
            <td class="pelatihan">${post.pelatihan}</td>
            <td class="keterangan">${post.keterangan}</td>
            <td><button class="button-edit" id="button-edit">Edit</button></td>
            <td><button class="button-delete" id="button-delete">Hapus</button></td>
        </tr>`;
    })
    tbody.innerHTML = data;
};

const postData = () => {
    return{
        name: Vname.value,
        kecamatan: Vkecamatan.value,
        pelatihan: Vpelatihan.value,
        keterangan: Vketerangan.value
    }
};

// get data
fetch("http://localhost:8888/sertificate/all")
.then(response => response.json())
.then(data => {
    const realData = data.data;
    console.log(data);
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

// edit & delete data
const submitEdit = document.querySelector(".submit-button");
const classData = document.querySelector('.t-body');
classData.addEventListener('click', (e) => {
    e.preventDefault();

    // delete data
    let id = e.target.parentElement.parentElement.dataset.id;
    if(e.target.classList.contains('button-delete')){
        fetch(`http://localhost:8888/sertificate/${id}`, {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(() => location.reload())
    }

    // edit data
    if(e.target.classList.contains('button-edit')){
        document.querySelector('.popup').style.display = 'flex';
    
        if(document.querySelector('.popup').style.display === 'flex') {
            document.querySelector('body').style.overflow = 'hidden';
            const parent = e.target.parentElement.parentElement;
            let name = parent.querySelector('.name').textContent;
            let kecamatan = parent.querySelector('.kecamatan').textContent;
            let pelatihan = parent.querySelector('.pelatihan').textContent;
            let keterangan = parent.querySelector('.keterangan').textContent;
    
            Vname.value = name;
            Vkecamatan.value = kecamatan;
            Vpelatihan.value = pelatihan;
            Vketerangan.value = keterangan;
        }
    }
    submitEdit.addEventListener('click', () => {
        fetch(`http://localhost:8888/sertificate/${id}`, {
            method: 'PUT',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify({
                name: Vname.value,
                kecamatan: Vkecamatan.value,
                pelatihan: Vpelatihan.value,
                keterangan: Vketerangan.value
            })
        })
        .then(response => response.json())
        .then(() => location.reload())
    })
});