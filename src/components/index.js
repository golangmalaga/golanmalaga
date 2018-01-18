import React, { Component } from 'react';
import { $, PeticionAjax } from '../lib/helpers';

 class App extends Component {
    constructor(...props){
        super();
        this.state = {
            user: []
        };
        this.handleSubmitLogin = this.handleSubmitLogin.bind(this);
        this.handleOnSubmitRegister = this.handleOnSubmitRegister.bind(this);
    }
    handleSubmitLogin(e){
        let formLogin = $('form-login'),
            email = $('email'),
            password = $('password'),
            // btnLogin = $('btnLogin'),
            mensajeLogin = $('mensaje-login');
            e.preventDefault();
            let obj = {
                email: email.value,
                password: password.value
            };

            PeticionAjax(formLogin.method, formLogin.action, JSON.stringify(obj))
                .then(respuesta => {
                    if (respuesta.status === 200) {
                        mensajeLogin.textContent ='Ingresaste';
                        sessionStorage.setItem('token', respuesta.response.token);
                        console.log(respuesta.response);
                    } else {
                        mensajeLogin.textContent = respuesta.response.message;
                        console.log(respuesta.response);
                    }
                })
                .catch(error => {
                    console.log(error);
                });
    }
    handleOnSubmitRegister (e) {
        let formRegister = $('form-register'),
            username = $('username'),
            emailRegister = $('email-register'),
            name = $('nombre'),
            passwordRegister = $('password-register'),
            confirmPassword= $('password-register2'),
            mensajeRegister = $('mensaje-register');

            e.preventDefault();
            let obj = {
                username: username.value,
                email: emailRegister.value,
                namefull: name.value,
                password: passwordRegister.value,
                confirmPassword: confirmPassword.value
            };
            PeticionAjax(formRegister.method, formRegister.action, JSON.stringify(obj))
                .then(respuesta => {
                    if (respuesta.status === 200) {
                        mensajeRegister.textContent = respuesta.response.message;
                    } else {
                        mensajeRegister.textContent = respuesta.response.message;
                    }
                })
                .catch(error => {
                    console.log(error);
                });
            formRegister.target.value = "";
    }
    render(){
        return(
            <div>
                <form action="/api/login" method="post" id="form-login" onSubmit={this.handleSubmitLogin}>
                    <input type="email" id="email" name="email" placeholder="Email"/>
                    <input type="password" id="password" name="password" placeholder="Password"/>
                    <button id="btnLogin">Login</button>
                </form>
                <div id="mensaje-login"></div>
                <form action="/api/users/" method="post" id="form-register" onSubmit={this.handleOnSubmitRegister}>
                    <input type="text" id="username" name="username" placeholder="Usuario"/>
                    <input type="email" id="email-register" name="email-register" placeholder="Email"/>
                    <input type="text" id="nombre" name="nombre" placeholder="nombre"/>
                    <input type="password" id="password-register" name="password-register" placeholder="Password"/>
                    <input type="password" id="password-register2" name="password-register2" placeholder="Confirmar Password"/>
                    <button id="btnLogin-register">Registrarte</button>
                </form>
                <div id="mensaje-register"></div>
            </div>
        )
    }
}
export default App