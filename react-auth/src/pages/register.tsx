import React,{SyntheticEvent, useState} from "react";
import { Navigate } from "react-router-dom";

const Register = () => {
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [redirect, setRedirect] = useState(false);

    const handleSubmit = async (e:SyntheticEvent) => {
        e.preventDefault();
        await fetch("http://localhost:8000/api/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                name,
                email,
                password,
            }),
        });
        setRedirect(true);
    }
    if (redirect){
        return <Navigate to="/login" />
    }
    
    return (
       <form>
                <h1 className="h3 mb-3 fw-normal">Please Register</h1>
                <div className="form-floating">
                    <input
                        type="name"
                        className="form-control"
                        placeholder="name@example.com"
                        required
                        onChange={(e) => setName(e.target.value)}
                    />
                    <label htmlFor="floatingInput">Name</label>{" "}
                </div>
                <div className="form-floating">
                    <input
                        type="email"
                        className="form-control"
                        placeholder="name@example.com"
                        required
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <label htmlFor="floatingInput">Email address</label>{" "}
                </div>
                <div className="form-floating">
                    <input
                        type="password"
                        className="form-control"
                        placeholder="Password"
                        required
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <label htmlFor="floatingPassword">Password</label>{" "}
                </div>
                <div className="form-check text-start my-3"></div>
                <button className="btn btn-primary w-100 py-2" type="submit" onClick={handleSubmit}>
                    Submit
                </button>
            </form>

    );
};
export default Register;