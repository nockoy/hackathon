import { useNavigate } from "react-router-dom";
const NotFound = () => {
  const navigate = useNavigate()
  return (
    <>
      <form>
        <h1>
          404 NotFound
        </h1>
        <h3>お探しのページは見つかりませんでした。</h3>
        <button  onClick={() => navigate('/')}>ホームに戻る</button>
      </form>
    </>
  );
};

export default NotFound;