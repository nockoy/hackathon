const Login = () => {
  return (
    <>
      <h1>ログインページ</h1>
      <form>
        <div>
          <label>メールアドレス</label>
          <input
            name="email"
            type="email"
          />
        </div>
        <div>
          <label>パスワード</label>
          <input
            name="password"
            type="password"
          />
        </div>
        <button>ログイン</button>
      </form>
    </>
  );
};

export default Login;