import React, { Component } from "react";
import axios from "axios";

const serverUrl = "http://115.85.181.25:8080";

class User{
    constructor(userId, name, address, email, level, password){
        this.userId = userId;
        this.name = name;
        this.address = address;
        this.email = email;
        this.level = level;
        this.password = password;
    }
}

class CreateUser extends Component{
    state = { userId:"", name:"", address:"", email:"", password:"", confirmPassword:""}

    handleChange = ({target : { name, value}}) => {
        this.setState({[name]: value});
    }

    handleSubmit = async (event) => {
        event.preventDefault()

        if (this.state.userId === ''){
            alert("INPUT ID!");
            return;
        }
        else if (this.state.name === ''){
            alert("INPUT NAME!");
            return;
        }
        else if (this.state.password === ''){
            alert("INPUT PASSWORD!");
            return;
        }    
        else if (this.state.password !== this.state.confirmPassword){
            alert("INPUT PASSWORD IS NOT EUQAL TO CONFIRM PASSWORD!");
            return;
        }

        var requestUrl = serverUrl + "/users";
        var user = new User(this.state.userId.toString(), this.state.name.toString(), this.state.address, this.state.email, 1, this.state.password.toString());

        axios
        .post(requestUrl, user)
        .then(response => {
            alert("Registration is done!");
            this.setState({userId:"", name:"",address:"",email:"", password:"", confirmPassword:"" });            
          })
        .catch(function (error) {
            if (error.response) {
                // 요청이 이루어졌으며 서버가 2xx의 범위를 벗어나는 상태 코드로 응답했습니다.
                if (error.response.data) {
                    alert(error.response.data["Detail"])
                    //alert(error.message)
                }
                else {
                    alert('Error: ' +  error.message)
                }
            }
            else if (error.request) {
              // 요청이 이루어 졌으나 응답을 받지 못했습니다.
              // `error.request`는 브라우저의 XMLHttpRequest 인스턴스 또는
              // Node.js의 http.ClientRequest 인스턴스입니다.
              //alert(error.request);
              alert(error.request)
            }
            else {
              // 오류를 발생시킨 요청을 설정하는 중에 문제가 발생했습니다.
              console.log('Error', error.message);
            }
          });
    }
      
    render() {

        return (
                <form onSubmit={this.handleSubmit} >
                    <h4>회원가입</h4>
                    <p></p>
                    <p> ID: <input type="text" name="userId" value={this.state.userId} onChange={this.handleChange}/> </p>
                    <p> NAME: <input type="text" name="name" value={this.state.name} onChange={this.handleChange}/></p>
                    <p> ADDRESS: <input type="text" name="address" value={this.state.address} onChange={this.handleChange}/></p>
                    <p> EMAIL: <input type="text" name="email" value={this.state.email} onChange={this.handleChange}/></p>
                    <p> PASSWORD: <input type="password" name="password" value={this.state.password} onChange={this.handleChange}/></p>
                    <p> CONFIRM PASSWORD: <input type="password" name="confirmPassword" value={this.state.confirmPassword} onChange={this.handleChange}/></p>
                    <button type="submit">가입</button>
                </form>);
    }
}

class SearchUser extends Component{
    state = { userId:"", userList: null }

    handleChange = ({target : { name, value}}) => {
        this.setState({[name]: value});
    }

    handleSubmit = async (event) => {
        event.preventDefault()

        var requestUrl = serverUrl + "/users";

        axios
        .get(requestUrl, {params: { userId: this.state.userId} })
        .then(response => {
            if (response.data["data"].length > 0)
            {
                this.setState({id: "",userList: <UserList data={response.data["data"]}/>})
            }
            else
            {
                this.setState({id: "",userList: <UserList data={null}/>})
            }                        
          })
          .catch(function (error) {
            if (error.response) {
                alert(error.message);
            }
          });
    }
      
    render() {
        return (
                <div>
                <form onSubmit={this.handleSubmit} >
                    <h4>회원검색</h4>
                    <p></p>
                    <p> USER ID: <input type="text" name="userId" value={this.state.userId} onChange={this.handleChange}/> </p>
                    <button type="submit">검색</button>
                </form>
                    <div>
                        {this.state.userList}
                    </div>
                </div>
            );
    }
}

class UserList extends Component{

    render() {        
        return (
            <table className="UserList"  border="1px solid black">
            <thead>
                <tr>
                    <th scope="col">USERID</th>
                    <th scope="col">NAME</th>
                    <th scope="col">ADDRESS</th>
                    <th scope="col">EMAIL</th>
                    <th scope="col">LEVEL</th>
                </tr>
            </thead>
            <tbody>
                {this.props.data && this.props.data.map((item) => 
                    <tr key={item.userId}>
                        <td>{item.userId}</td>
                        <td>{item.name}</td>
                        <td>{item.address}</td>
                        <td>{item.email}</td>
                        <td>{item.level}</td>
                    </tr>)}
            </tbody>
        </table>
        );
    }
}

/*
class UpdateUser extends Component{
    state = { userId:"", userInfo: null }

    handleChange = ({target : { name, value}}) => {
        this.setState({[name]: value});
    }

    handleSubmit = async (event) => {
        event.preventDefault()

        var requestUrl = serverUrl + "/users";

        axios
        .get(requestUrl, {params: { userId: this.state.userId} })
        .then(response => {            
            if (response.data.length === 0)
            {
                this.setState({id: "",userInfo: <UserInfo data={null}/>})
            }
            else
            {
                this.setState({id: "",userInfo: <UserInfo data={response.data["data"]}/>})
            }
          })
          .catch(function (error) {
            if (error.response) {
                alert(error.message);
            }
          });
    }
      
    render() {
        return (
                <div>
                <form onSubmit={this.handleSubmit} >
                    <h4>회원검색</h4>
                    <p></p>
                    <p> USER ID: <input type="text" name="userId" value={this.state.userId} onChange={this.handleChange}/> </p>
                    <button type="submit">검색</button>
                </form>
                    <div>
                        {this.state.userInfo}
                    </div>
                </div>
            );
    }
}

class UserInfo extends Component{
    state = { userId:"", name:"", address:"", email:"", password:"", confirmPassword:""}

    handleChange = ({target : { name, value}}) => {
        this.setState({[name]: value});
    }

    handleSubmit = async (event) => {
        event.preventDefault()

        if (this.state.userId === ''){
            alert("INPUT ID!");
            return;
        }
        else if (this.state.name === ''){
            alert("INPUT NAME!");
            return;
        }
        else if (this.state.password === ''){
            alert("INPUT PASSWORD!");
            return;
        }    
        else if (this.state.password !== this.state.confirmPassword){
            alert("INPUT PASSWORD IS NOT EUQAL TO CONFIRM PASSWORD!");
            return;
        }

        var requestUrl = serverUrl + "/users";
        var user = new User(this.state.userId.toString(), this.state.name.toString(), this.state.address, this.state.email, 1, this.state.password.toString());

        axios
        .post(requestUrl, user)
        .then(response => {
            alert("Registration is done!");
            this.setState({userId:"", name:"",address:"",email:"", password:"", confirmPassword:"" });            
          })
        .catch(function (error) {
            if (error.response) {
                // 요청이 이루어졌으며 서버가 2xx의 범위를 벗어나는 상태 코드로 응답했습니다.
                if (error.response.data) {
                    alert(error.response.data["Detail"])
                    //alert(error.message)
                }
                else {
                    alert('Error: ' +  error.message)
                }
            }
            else if (error.request) {
              // 요청이 이루어 졌으나 응답을 받지 못했습니다.
              // `error.request`는 브라우저의 XMLHttpRequest 인스턴스 또는
              // Node.js의 http.ClientRequest 인스턴스입니다.
              //alert(error.request);
              alert(error.request)
            }
            else {
              // 오류를 발생시킨 요청을 설정하는 중에 문제가 발생했습니다.
              console.log('Error', error.message);
            }
          });
    }
      
    render() {

        return (
                <form onSubmit={this.handleSubmit} >
                    <h4>회원가입</h4>
                    <p></p>
                    <p> ID: <input type="text" name="userId" value={this.state.userId} onChange={this.handleChange}/> </p>
                    <p> NAME: <input type="text" name="name" value={this.state.name} onChange={this.handleChange}/></p>
                    <p> ADDRESS: <input type="text" name="address" value={this.state.address} onChange={this.handleChange}/></p>
                    <p> EMAIL: <input type="text" name="email" value={this.state.email} onChange={this.handleChange}/></p>
                    <p> PASSWORD: <input type="password" name="password" value={this.state.password} onChange={this.handleChange}/></p>
                    <p> CONFIRM PASSWORD: <input type="password" name="confirmPassword" value={this.state.confirmPassword} onChange={this.handleChange}/></p>
                    <button type="submit">가입</button>
                </form>);
    }
}
*/
export {CreateUser, SearchUser};
