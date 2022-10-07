import axios from "axios";
import React, { Component } from "react";
import { Link } from 'react-router-dom';
import config from "./config.json";
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { Password } from 'primereact/password';
import "primereact/resources/themes/lara-light-indigo/theme.css";  //theme
import "primereact/resources/primereact.min.css";                  //core css
import "primeicons/primeicons.css";                                //icons



class Login extends Component {
    state = {id:"", password:""}

    handleChange = ({target : { name, value}}) => {
        this.setState({[name]: value});
    }

    OnSignInClicked = async (event) => {
        event.preventDefault()
        alert("Ddd")

        var requestUrl = config.SERVER_URL + "/user/login"
        await axios.post(requestUrl, { Id:this.state.id, password:this.state.password})
    }

    render() {
        return (
            <div>
                <form>
                    <h4>Humor Blog</h4>
                    <p></p>
                    <p> ID: <InputText name="id" value={this.state.id} onChange={this.handleChange} /> </p>
                    <p> PASSWORD: <Password name="password" value={this.state.password} onChange={this.handleChange}/> </p>
                    <Button label="Sign In" type="submit" onSubmit={this.OnSignInClicked}></Button>
                    <Link to="/createUser">
                        <Button label="Sign Up"></Button>
                    </Link>
                </form>
            </div>
        );
    }
}

export { Login };