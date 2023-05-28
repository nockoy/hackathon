import Sidebar from "../components/Sidebar";
import Topbar from "../components/Topbar";
import Header from "../components/Header";
import MessageField from "../components/MessageField";
import SendBox from "../components/Sendbox";

export default function Mypage() {
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
          <Header />
        </div>
        <div id="itemC">
          <MessageField />
        </div>
        <div id="itemD">
          <SendBox/>
        </div>
      </div>
    </div>
  );
}