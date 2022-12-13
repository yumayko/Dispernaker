const V2name = document.getElementById("namaEdit");
const V2kecamatan = document.getElementById("kecamatanEdit");
const V2pelatihan = document.getElementById("pelatihanEdit");
const V2keterangan = document.getElementById("keteranganEdit");
const V2sertifikasi = document.getElementById("sertifikasiEdit");
const submitEdit = document.querySelector(".edit-button");

// DISPLAY DATA
const url1 = fetch('http://54.226.168.28/minat/all');
const url2 = fetch('http://54.226.168.28/pelatihan/all');
const url3 = fetch('http://54.226.168.28/sertifikasi/calon/all');
const url4 = fetch('http://54.226.168.28/sertifikasi/peserta/all');

Promise.all([url1, url2, url3, url4])
.then(responses => {
    return Promise.all(responses.map(response => response.json()));
}).then(([calonPel, pesertaPel, calonSert, pesertaSert]) => {
    GetData1(calonPel.data);
    GetData2(pesertaPel.data);
    GetData3(calonSert.data);
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
            <td class="name">${post.nama}</td>
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
            <td class="name">${post.nama}</td>
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
            <td class="name">${post.nama}</td>
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
            <td class="name">${post.nama}</td>
            <td class="kecamatan">${post.kecamatan}</td>
            <td class="pelatihan">${post.pelatihan}</td>
            <td class="sertifikasi">${post.sertifikasi}</td>
            <td class="keterangan">${post.keterangan}</td>
            <td><button class="pesertaSertEdit" id="button-edit">Edit</button></td>
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
const Vsertifikasi = document.getElementById("sertifikasi");
const postData1 = () => {
    return{
        nama: Vname.value,
        kecamatan: Vkecamatan.value,
        pelatihan: Vpelatihan.value,
        keterangan: Vketerangan.value
    }
};
const postData2 = () => {
    return{
        nama: Vname.value,
        kecamatan: Vkecamatan.value,
        pelatihan: Vpelatihan.value,
        keterangan: Vketerangan.value,
        sertifikasi: Vsertifikasi.value
    }
};

const addForm = document.querySelector('.popup-body');
addForm.addEventListener('click', async (e) => {
    e.preventDefault();
    let payload = postData1();
    let payload2 = postData2();
    if(e.target.id==='submit-calonPel'){
        await postMinat(payload);
    }
    if(e.target.id==='submit-pesertaPel'){
        await postPelatihan(payload);
    }
    if(e.target.id==='submit-calonSert'){
        await postcalonSert(payload);
    }
    if(e.target.id==='submit-pesertaSert'){
        await postpesertaSert(payload2);
    }
});

function postMinat(postData){
    return fetch("http://54.226.168.28/minat", {
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
function postPelatihan(postData){
    return fetch("http://54.226.168.28/pelatihan", {
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
    return fetch("http://54.226.168.28/sertifikasi/calon", {
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
    return fetch("http://54.226.168.28/sertifikasi/peserta", {
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
const classData = document.querySelectorAll('.t-body');
for (let i = 0; i < classData.length; i++) {
    classData[i].addEventListener('click', (e) => {
        let id = e.target.parentElement.parentElement.dataset.id;
        // delete data
        if(e.target.classList.contains('button-delete')){
            if(e.target.parentElement.parentElement.parentElement.id==='tbody1'){
                deleteMinat(id);
            }
            if(e.target.parentElement.parentElement.parentElement.id==='tbody2'){
                deletePelatihan(id);
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
            document.querySelector('.SertifikasiEdit').style.display = 'none';
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
        if(e.target.classList.contains('pesertaSertEdit')){
            const popupEdit = document.querySelector('.popupEdit').style.display = 'flex';
            document.querySelector('.SertifikasiEdit').style.display = 'block';
            if(popupEdit === 'flex') {
                document.querySelector('body').style.overflow = 'hidden';
                const parent = e.target.parentElement.parentElement;
                let name = parent.querySelector('.name').textContent;
                let kecamatan = parent.querySelector('.kecamatan').textContent;
                let pelatihan = parent.querySelector('.pelatihan').textContent;
                let sertifikasi = parent.querySelector('.sertifikasi').textContent;
                let keterangan = parent.querySelector('.keterangan').textContent;
        
                V2name.value = name;
                V2kecamatan.value = kecamatan;
                V2pelatihan.value = pelatihan;
                V2sertifikasi.value = sertifikasi;
                V2keterangan.value = keterangan;
            }
        }
        submitEdit.addEventListener('click', (a) => {
            a.preventDefault();
            if(e.target.parentElement.parentElement.parentElement.id==='tbody1'){
                editMinat(id);
            }
            if(e.target.parentElement.parentElement.parentElement.id==='tbody2'){
                editPelatihan(id);
            }
            if(e.target.parentElement.parentElement.parentElement.id==='tbody3'){
                editcalonSert(id);
            }
            if(e.target.parentElement.parentElement.parentElement.id==='tbody4'){
                editpesertaSert(id);
            }
        })
    });
}

//fetch delete data
function deleteMinat(idDelete){
    return fetch("http://54.226.168.28/minat/" + (idDelete), {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(() => location.reload())
}
function deletePelatihan(idDelete){
    return fetch("http://54.226.168.28/pelatihan/" + (idDelete), {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(() => location.reload())
}
function deletecalonSert(idDelete){
    return fetch("http://54.226.168.28/sertifikasi/calon/" + (idDelete), {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(() => location.reload())
}
function deletepesertaSert(idDelete){
    return fetch("http://54.226.168.28/sertifikasi/peserta/" + (idDelete), {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(() => location.reload())
}

//fetch edit data
function editMinat(idEdit){
    return fetch("http://54.226.168.28/minat/" + (idEdit), {
            method: 'PUT',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify({
                nama: V2name.value,
                kecamatan: V2kecamatan.value,
                pelatihan: V2pelatihan.value,
                keterangan: V2keterangan.value
            })
        })
        .then(response => response.json())
        .then(data => location.reload())
}
function editPelatihan(idEdit){
    return fetch("http://54.226.168.28/pelatihan/" + (idEdit), {
            method: 'PUT',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify({
                nama: V2name.value,
                kecamatan: V2kecamatan.value,
                pelatihan: V2pelatihan.value,
                keterangan: V2keterangan.value
            })
        })
        .then(response => response.json())
        .then(data => location.reload())
}
function editcalonSert(idEdit){
    return fetch("http://54.226.168.28/sertifikasi/calon/" + (idEdit), {
            method: 'PUT',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify({
                nama: V2name.value,
                kecamatan: V2kecamatan.value,
                pelatihan: V2pelatihan.value,
                keterangan: V2keterangan.value
            })
        })
        .then(response => response.json())
        .then(data => location.reload())
}
function editpesertaSert(idEdit){
    return fetch("http://54.226.168.28/sertifikasi/peserta/" + (idEdit), {
            method: 'PUT',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify({
                nama: V2name.value,
                kecamatan: V2kecamatan.value,
                pelatihan: V2pelatihan.value,
                sertifikasi: V2sertifikasi.value,
                keterangan: V2keterangan.value
            })
        })
        .then(response => response.json())
        .then(data => location.reload())
}