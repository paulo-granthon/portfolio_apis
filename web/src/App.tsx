import Portfolio from "./pages/portfolio";
import { Routes, Route, BrowserRouter } from "react-router-dom";

export default function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/portfolio/:userId" element={<Portfolio />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}
