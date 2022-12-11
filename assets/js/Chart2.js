async function chartKec() {
  await getMinatData();
  await getPelatihanData();
  await getcalonSertData();
  await getpesertaSertData();

  const ctx = document.getElementById('chartKec');

  const chart = new Chart(ctx, {
      // The type of chart we want to create
      type: 'bar',

      // The data for our dataset
      data: {
          labels: NewLabels,
          datasets: [
            {
              label: 'Peserta Minat',
              data: dataKecMinat,
              backgroundColor: 'rgb(26, 163, 255, 0.7)',
              borderWidth: 1
            },
            {
              label: 'Peserta Pelatihan',
              data: dataKecPelatihan,
              backgroundColor: 'rgb(255, 99, 132, 0.7)',
              borderWidth: 1
            },
            {
              label: 'Calon Sertifikasi',
              data: dataKecSertC,
              backgroundColor: 'rgb(255, 205, 86, 0.7)',
              borderWidth: 1
            },
            {
              label: 'Peserta Sertifikasi',
              data: dataKecSertP,
              backgroundColor: 'rgb(210, 77, 255, 0.7)',
              borderWidth: 1
            }
          ]
      },

      // Configuration options go here
      options: {
          tooltips: {
              mode: 'index'
          }
        }
  });
}
async function chartPel() {
  await getMinatData();
  await getPelatihanData();
  await getcalonSertData();
  await getpesertaSertData();

  const ctx = document.getElementById('chartPel');

  const chart = new Chart(ctx, {
      // The type of chart we want to create
      type: 'bar',

      // The data for our dataset
      data: {
          labels: LabelPel,
          datasets: [
            {
              label: 'Peserta Minat',
              data: dataPelMinat,
              backgroundColor: 'rgb(26, 163, 255, 0.7)',
              borderWidth: 1
            },
            {
              label: 'Peserta Pelatihan',
              data: dataPelPelatihan,
              backgroundColor: 'rgb(255, 99, 132, 0.7)',
              borderWidth: 1
            },
            {
              label: 'Calon Sertifikasi',
              data: dataPelSertC,
              backgroundColor: 'rgb(255, 205, 86, 0.7)',
              borderWidth: 1
            },
            {
              label: 'Peserta Sertifikasi',
              data: dataPelSertP,
              backgroundColor: 'rgb(210, 77, 255, 0.7)',
              borderWidth: 1
            }
          ]
      },

      // Configuration options go here
      options: {
          tooltips: {
              mode: 'index'
          }
        }
  });
}

chartPel();
chartKec();

//Fetch Data from API
let elementCnt1=[], NewLabels=[], dataKecMinat=[], dataKecPelatihan=[], dataKecSertC=[], dataKecSertP=[];
let elementCnt2=[], LabelPel=[], dataPelMinat=[], dataPelPelatihan=[], dataPelSertC=[], dataPelSertP=[];
function getMinatData(){
  return fetch("http://54.226.168.28/minat/all")
  .then(response => response.json())
  .then(res => {
      const data = res.data;
      //kecamatan
      const labelsKec = data.map((x) => x.kecamatan);
      labelsKec.forEach(val => elementCnt1[val] = (elementCnt1[val] || 0) + 1);
      NewLabels = (Object.keys(elementCnt1));
      dataKecMinat = (Object.values(elementCnt1));

      //Pelatihan
      const labelsPel = data.map((x) => x.pelatihan);
      labelsPel.forEach(val => elementCnt2[val] = (elementCnt2[val] || 0) + 1);
      LabelPel = (Object.keys(elementCnt2));
      console.log(LabelPel);
      dataPelMinat = (Object.values(elementCnt2));
  })
}
function getPelatihanData(){
  return fetch("http://54.226.168.28/pelatihan/all")
  .then(response => response.json())
  .then(res => {
      const data = res.data;
      //kecamatan
      const labels = data.map((x) => x.kecamatan);
      labels.forEach(val => elementCnt1[val] = (elementCnt1[val] || 0) + 1);
      dataKecPelatihan = (Object.values(elementCnt1));

      //Pelatihan
      const labelsPel = data.map((x) => x.pelatihan);
      labelsPel.forEach(val => elementCnt2[val] = (elementCnt2[val] || 0) + 1);
      dataPelPelatihan = (Object.values(elementCnt2));
  })
}
function getcalonSertData(){
  return fetch("http://54.226.168.28/sertifikasi/calon/all")
  .then(response => response.json())
  .then(res => {
      const data = res.data;
      //Pelatihan
      const labels = data.map((x) => x.kecamatan);
      labels.forEach(val => elementCnt1[val] = (elementCnt1[val] || 0) + 1);
      dataKecSertC = (Object.values(elementCnt1));

      //Pelatihan
      const labelsPel = data.map((x) => x.pelatihan);
      labelsPel.forEach(val => elementCnt2[val] = (elementCnt2[val] || 0) + 1);
      dataPelSertC = (Object.values(elementCnt2));
  })
}
function getpesertaSertData(){
  return fetch("http://54.226.168.28/sertifikasi/peserta/all")
  .then(response => response.json())
  .then(res => {
      const data = res.data;
      //Pelatihan
      const labels = data.map((x) => x.kecamatan);
      labels.forEach(val => elementCnt1[val] = (elementCnt1[val] || 0) + 1);
      dataKecSertP = (Object.values(elementCnt1));

      //Pelatihan
      const labelsPel = data.map((x) => x.pelatihan);
      labelsPel.forEach(val => elementCnt2[val] = (elementCnt2[val] || 0) + 1);
      dataPelSertP = (Object.values(elementCnt2));
  })
}