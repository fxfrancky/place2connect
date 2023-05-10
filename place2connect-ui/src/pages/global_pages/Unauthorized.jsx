import { useNavigate } from "react-router-dom"

const Unauthorized = () => {
    const navigate = useNavigate();

    // const goBack = () => navigate(-1);
    const goBack = () => navigate("/");

    return (
        <section>
            <h1>Unauthorized</h1>
            <br />
            <p>You do not have access to the requested page.</p>
            <div className="flexGrow">
                {/* <button onClick={goBack}>Go Back</button> */}
                <button onClick={goBack}>Home Page</button>

            </div>
        </section>
    )
}

export default Unauthorized