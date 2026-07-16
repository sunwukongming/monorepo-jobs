import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { APP_NAME } from "@jobs/shared";
import App from "./App";
import "./styles.css";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <App title={APP_NAME} />
  </StrictMode>,
);
