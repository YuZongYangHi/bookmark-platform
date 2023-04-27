import {useState} from "react";

export default () => {
  const [menuRouterList, setMenuRouterList] = useState([]);
  return {
    menuRouterList, setMenuRouterList
  }
}
