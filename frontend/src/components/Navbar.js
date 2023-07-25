import React from 'react';
import { Layout, Menu } from 'antd';
const { Header } = Layout;

const items1 = [
    {key: '1', label: 'logout'}
]

const Navbar = (props) => {

    return (
        <Header
            style={{
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'flex-end'
            }}
          >
            <div className="demo-logo" />
            <Menu theme="dark" mode="horizontal" items={items1} onClick={props.logout}/>
          </Header>
    );
}

export default Navbar