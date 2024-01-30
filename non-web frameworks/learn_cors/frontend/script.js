// URL backend Anda
const backendURL = 'http://localhost:9000';
const testToken = 'your_test_token_here'; // Gantilah dengan token yang sesuai

// Membuat permintaan GET menggunakan fetch dengan header Authorization
fetch(`${backendURL}/index`, {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${testToken}`, // Menambahkan token ke header Authorization
  },
})
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
    return response.text();
  })
  .then(data => {
    console.log(data);
  })
  .catch(error => {
    console.error('Fetch error:', error);
  });
