const Vname = document.getElementById("nama");
const Vkecamatan = document.getElementById("kecamatan");
const Vpelatihan = document.getElementById("pelatihan");
const Vketerangan = document.getElementById("keterangan");
const V2name = document.getElementById("namaEdit");
const V2kecamatan = document.getElementById("kecamatanEdit");
const V2pelatihan = document.getElementById("pelatihanEdit");
const V2keterangan = document.getElementById("keteranganEdit");
const addForm = document.querySelector('.popup-body');
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
const postData = () => {
    return{
        name: Vname.value,
        kecamatan: Vkecamatan.value,
        pelatihan: Vpelatihan.value,
        keterangan: Vketerangan.value
    }
};

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
        document.querySelector('.popupEdit').style.display = 'flex';
    
        if(document.querySelector('.popupEdit').style.display === 'flex') {
            document.querySelector('body').style.overflow = 'hidden';
            console.log(id)
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
    console.log(V2name.value);
    console.log(V2kecamatan.value);
    console.log(V2pelatihan.value);
    console.log(V2keterangan.value);
    submitEdit.addEventListener('click', (a) => {
        a.preventDefault();
        fetch(`http://localhost:8888/sertificate/${id}`,{
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
    })
});