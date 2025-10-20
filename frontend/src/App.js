import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import LoginForm from "./components/LoginForm";
import StyleDivForm from "./components/UI/StyleDivForm";
import Account from "./components/Account";
import CallBox from "./components/CallBox";

function App() {
  return (
    <Router>
      <div className="App">
        <Routes>
          <Route
          path="/"
          element={
            <StyleDivForm>
              <LoginForm />
            </StyleDivForm>
          }/>
          <Route path="account" element={<Account />}/>
          <Route path="/callBox" element={<CallBox />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
