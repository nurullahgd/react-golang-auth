import React from 'react';
import { Link } from 'react-router-dom';

const Nav = (props: { name: string, setName:(name:string)=>void }) => {
    const handleLogOut = async (e: React.SyntheticEvent) => {
        e.preventDefault();
        await fetch("http://localhost:8000/api/logout", {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
        });
        window.location.href = "/login";
        props.setName("");
    }

    let menu;
        
    if (props.name === "") {
        menu = (<ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item">
                {" "}
                <Link to="/login" className="nav-link" >
                    Login
                </Link>{" "}
            </li>
            <li className="nav-item">
                {" "}
                <Link to="/register" className="nav-link" >
                    Register
                </Link>{" "}
            </li>
        </ul>)
    } else {
        menu = (<ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item active nav-link">
                {props.name}
            </li>
            <li className="nav-item">
                {" "}
                <Link to="/logout" className="nav-link" onClick={handleLogOut} >
                    Logout
                </Link>{" "}
            </li>
        </ul>)
    }
    return (
        <div>
            <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
                <div className="container-fluid">
                    <Link to="/" className="navbar-brand" >
                        Home
                    </Link>

                    <div >
                        {menu}

                    </div>{" "}
                </div>{" "}
            </nav>{" "}
        </div>

    );
};
export default Nav;