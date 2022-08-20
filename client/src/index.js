import React from 'react';
import ReactDOM from 'react-dom/client';
import { Routes, Route } from 'react-router-dom';
import { BrowserRouter as Router } from 'react-router-dom';
import { Link } from 'react-router-dom';

// eslint-disable-next-line
import { CreateUser, SearchUser } from './User';

function Welcome() {

  return (
    <div className='Welcome'>
      <h3>유머 게시판</h3>
       <div>
        <Link to="/createUser">
          <button> 회원가입</button>
        </Link>
        <Link to="/searchUser">
          <button>회원검색</button>
        </Link>
        <Link to="/searchUser">
          <button>회원정보수정</button>
        </Link>
        <Link to="/searchUser">
          <button>글쓰기</button>
        </Link>
        <Link to="/searchUser">
          <button>내가 쓴 글 조회</button>
        </Link>
          <Routes>
            <Route path="/createUser" element={<CreateUser />} />
            <Route path="/searchUser" element={<SearchUser />} />
          </Routes>
        </div>
    </div>);
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <Router>
      <Welcome />
    </Router>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
//reportWebVitals();
