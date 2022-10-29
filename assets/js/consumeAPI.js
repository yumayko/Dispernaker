fetch("http://localhost:8888/health")
.then(response => response.json())
.then(data => console.log(data));