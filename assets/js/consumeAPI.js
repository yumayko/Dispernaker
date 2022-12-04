const V2name = document.getElementById("namaEdit");
const V2kecamatan = document.getElementById("kecamatanEdit");
const V2pelatihan = document.getElementById("pelatihanEdit");
const V2keterangan = document.getElementById("keteranganEdit");
const submitEdit = document.querySelector(".edit-button");

// DISPLAY DATA
const url1 = fetch('http://localhost:8888/training/calon/all');
const url2 = fetch('http://localhost:8888/training/peserta/all');
const url3 = fetch('http://localhost:8888/sertificate/calon/all');
const url4 = fetch('http://localhost:8888/sertificate/peserta/all');

Promise.all([url1, url2, url3, url4])
.then(responses => {
    return Promise.all(responses.map(response => response.json()));
}).then(([calonPel, pesertaPel, calonSert, pesertaSert]) => {
    console.log(calonPel.data);
    GetData1(calonPel.data);
    console.log(pesertaPel.data);
    GetData2(pesertaPel.data);
    console.log(calonSert.data);
    GetData3(calonSert.data);
    console.log(pesertaSert.data);
    GetData4(pesertaSert.data);
}).catch(err => {
    console.log(err);
})

const GetData1 = (posts) =>{
    let data = "";
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
    document.getElementById("tbody1").innerHTML = data;
    $(document).ready(function () {
        $('table.display').DataTable();
    });
};
const GetData2 = (posts) =>{
    let data = "";
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
    document.getElementById("tbody2").innerHTML = data;
    $(document).ready(function () {
        $('table.display').DataTable();
    });
};
const GetData3 = (posts) =>{
    let data = "";
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
    document.getElementById("tbody3").innerHTML = data;
    $(document).ready(function () {
        $('table.display').DataTable();
    });
};
const GetData4 = (posts) =>{
    let data = "";
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
    document.getElementById("tbody4").innerHTML = data;
    $(document).ready(function () {
        $('table.display').DataTable();
    });
};

// post data (create data)
const Vname = document.getElementById("nama");
const Vkecamatan = document.getElementById("kecamatan");
const Vpelatihan = document.getElementById("pelatihan");
const Vketerangan = document.getElementById("keterangan");
const postData = () => {
    return{
        name: Vname.value,
        kecamatan: Vkecamatan.value,
        pelatihan: Vpelatihan.value,
        keterangan: Vketerangan.value
    }
};

const addForm = document.querySelector('.popup-body');
addForm.addEventListener('click', async (e) => {
    e.preventDefault();
    let payload = postData();
    if(e.target.id==='submit-calonPel'){
        await postcalonPel(payload);
    }
    if(e.target.id==='submit-pesertaPel'){
        await postpesertaPel(payload);
    }
    if(e.target.id==='submit-calonSert'){
        await postcalonSert(payload);
    }
    if(e.target.id==='submit-pesertaSert'){
        await postpesertaSert(payload);
    }
});

function postcalonPel(postData){
    return fetch("http://localhost:8888/training/calon", {
            method: 'POST',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify(postData)
        })
        .then(response => response.json())
        .then(data => {
            const dataArr = [];
            dataArr.push(data.data);
            GetData1(dataArr);
            location.reload();
        })
}
function postpesertaPel(postData){
    return fetch("http://localhost:8888/training/peserta", {
            method: 'POST',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify(postData)
        })
        .then(response => response.json())
        .then(data => {
            const dataArr = [];
            dataArr.push(data.data);
            GetData2(dataArr);
            location.reload();
        })
}
function postcalonSert(postData){
    return fetch("http://localhost:8888/sertificate/calon", {
            method: 'POST',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify(postData)
        })
        .then(response => response.json())
        .then(data => {
            const dataArr = [];
            dataArr.push(data.data);
            GetData3(dataArr);
            location.reload();
        })
}
function postpesertaSert(postData){
    return fetch("http://localhost:8888/sertificate/peserta", {
            method: 'POST',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify(postData)
        })
        .then(response => response.json())
        .then(data => {
            const dataArr = [];
            dataArr.push(data.data);
            GetData4(dataArr);
            location.reload();
        })
}

// edit & delete data
const classData = document.querySelector('.t-body');
document.addEventListener('click', (e) => {
    let id = e.target.parentElement.parentElement.dataset.id;

    // delete data
    if(e.target.classList.contains('button-delete')){
        if(e.target.parentElement.parentElement.parentElement.id==='tbody1'){
            deletecalonPel(id);
        }
        if(e.target.parentElement.parentElement.parentElement.id==='tbody2'){
            deletepesertaPel(id);
        }
        if(e.target.parentElement.parentElement.parentElement.id==='tbody3'){
            deletecalonSert(id);
        }
        if(e.target.parentElement.parentElement.parentElement.id==='tbody4'){
            deletepesertaSert(id);
        }
    }

    // edit data
    if(e.target.classList.contains('button-edit')){
        const popupEdit = document.querySelector('.popupEdit').style.display = 'flex';
        if(popupEdit === 'flex') {
            document.querySelector('body').style.overflow = 'hidden';
            const parent = e.target.parentElement.parentElement;
            let name = parent.querySelector('.name').textContent;
            let kecamatan = parent.querySelector('.kecamatan').textContent;
            let pelatihan = parent.querySelector('.pelatihan').textContent;
            let keterangan = parent.querySelector('.keterangan').textContent;
    
            V2name.value = name;
            V2kecamatan.value = kecamatan;
            V2pelatihan.value = pelatihan;
            V2keterangan.value = keterangan;
        }
    }
    submitEdit.addEventListener('click', (a) => {
        a.preventDefault();
        if(e.target.parentElement.parentElement.parentElement.id==='tbody1'){
            editcalonPel(id);
        }
        if(e.target.parentElement.parentElement.parentElement.id==='tbody2'){
            editpesertaPel(id);
        }
        if(e.target.parentElement.parentElement.parentElement.id==='tbody3'){
            editcalonSert(id);
        }
        if(e.target.parentElement.parentElement.parentElement.id==='tbody4'){
            editpesertaSert(id);
        }
    })
});

//fetch delete data
function deletecalonPel(idDelete){
    return fetch("http://localhost:8888/training/calon/" + (idDelete), {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(() => location.reload())
}
function deletepesertaPel(idDelete){
    return fetch("http://localhost:8888/training/peserta/" + (idDelete), {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(() => location.reload())
}
function deletecalonSert(idDelete){
    return fetch("http://localhost:8888/sertificate/calon/" + (idDelete), {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(() => location.reload())
}
function deletepesertaSert(idDelete){
    return fetch("http://localhost:8888/sertificate/peserta/" + (idDelete), {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(() => location.reload())
}

//fetch edit data
function editcalonPel(idEdit){
    return fetch("http://localhost:8888/training/calon/" + (idEdit), {
            method: 'PUT',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify({
                name: V2name.value,
                kecamatan: V2kecamatan.value,
                pelatihan: V2pelatihan.value,
                keterangan: V2keterangan.value
            })
        })
        .then(response => response.json())
        .then(data => location.reload())
}
function editpesertaPel(idEdit){
    return fetch("http://localhost:8888/training/peserta/" + (idEdit), {
            method: 'PUT',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify({
                name: V2name.value,
                kecamatan: V2kecamatan.value,
                pelatihan: V2pelatihan.value,
                keterangan: V2keterangan.value
            })
        })
        .then(response => response.json())
        .then(data => location.reload())
}
function editcalonSert(idEdit){
    return fetch("http://localhost:8888/sertificate/calon/" + (idEdit), {
            method: 'PUT',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify({
                name: V2name.value,
                kecamatan: V2kecamatan.value,
                pelatihan: V2pelatihan.value,
                keterangan: V2keterangan.value
            })
        })
        .then(response => response.json())
        .then(data => location.reload())
}
function editpesertaSert(idEdit){
    return fetch("http://localhost:8888/sertificate/peserta/" + (idEdit), {
            method: 'PUT',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify({
                name: V2name.value,
                kecamatan: V2kecamatan.value,
                pelatihan: V2pelatihan.value,
                keterangan: V2keterangan.value
            })
        })
        .then(response => response.json())
        .then(data => location.reload())
}