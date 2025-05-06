import React,{useState} from "react";
import { Navigate } from "react-router-dom";

const Login = (props:{setName:(name:string)=>void}) => {
        const [email, setEmail] = useState("");
        const [password, setPassword] = useState("");
        const [redirect, setRedirect] = useState(false);

        const handleSubmit = async (e:React.SyntheticEvent) => {
            e.preventDefault();
            const response= await fetch("http://localhost:8000/api/login", {
                method: "POST",
                credentials: "include",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    email,
                    password,
                }),
            })

            const content =await response.json();
            setRedirect(true);
            props.setName(content.name);
        }
        if (redirect){
            return <Navigate to="/" />
        }
    return (
        <div>
            <form>
                <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
                <div className="form-floating">
                    <input
                        type="email"
                        className="form-control"
                        placeholder="name@example.com"
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                    <label htmlFor="floatingInput">Email address</label>{" "}
                </div>{" "}
                <div className="form-floating">
                    <input
                        type="password"
                        className="form-control"
                        placeholder="Password"
                        required
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <label htmlFor="floatingPassword">Password</label>{" "}
                </div>{" "}
                <div className="form-check text-start my-3"></div>
                <button className="btn btn-primary w-100 py-2" type="submit" onClick={handleSubmit}>
                    Sign in
                </button>
            </form>{" "}
        </div>

    );
};
export default Login;