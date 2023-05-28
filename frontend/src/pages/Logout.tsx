import { useNavigate } from "react-router-dom";
import Sidebar from "../components/Sidebar";
import Topbar from "../components/Topbar";

const Logout = () => {
    const navigate = useNavigate()
    return (
      <div className="App">
      <div id="GridContainer">
        <div id="Top">
          <Topbar />
        </div>
        <div id="itemA">
          <Sidebar />
        </div>
        <div id="itemB">
          B
        </div>
        <div id="itemC">
          C
          <h1>ログアウトしますか？</h1>
          <button  onClick={() => navigate('/')}>ホームに戻る</button>
          <button>ログアウト</button>
        </div>
        <div id="itemD">
          D
        </div>
      </div>
    </div>
      
    /*
      <Sidebar />
      
        <h1>ログアウトしますか？</h1>
        <button  onClick={() => navigate('/1/')}>ホームに戻る</button>

        <button>ログアウト</button>
    */

    );
  };
  
  export default Logout;

        