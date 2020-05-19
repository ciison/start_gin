let axios = require('axios');

let indexGet = function () {
    axios.get('http://localhost:8080/index').then((res) => console.log(res.data),
        (err) => console.log(err),)
}

indexGet();