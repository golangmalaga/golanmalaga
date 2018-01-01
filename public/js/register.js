let formRegister = $('form-register'),
    username = $('username'),
    emailRegister = $('email-register'),
    name = $('nombre'),
    passwordRegister = $('password-register'),
    confirmPassword= $('password-register2'),
    mensajeRegister = $('mensaje-register');

formRegister.addEventListener('submit', e => {
    e.preventDefault();
    let obj = {
        username: username.value,
        email: emailRegister.value,
        namefull: name.value,
        password: passwordRegister.value,
        confirmPassword: confirmPassword.value
    };
    peticionAjax(formRegister.method, formRegister.action, JSON.stringify(obj))
        .then(respuesta => {
            if (respuesta.status === 200) {
                mensajeLogin.textContent = respuesta.response.message;
            } else {
                mensajeLogin.textContent = respuesta.response.message;
            }
        })
        .catch(error => {
            console.log(error);
        });
});