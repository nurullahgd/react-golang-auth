import React from "react";

const Home = (props:{name:string}) => {
    return (
        <div>
            {props.name ? `Hi ${props.name}` : 'Please Login'}
        </div>

    );
};
export default Home;