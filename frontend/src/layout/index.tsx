import React, { useEffect, useState } from 'react';
import {
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    UploadOutlined,
    UserOutlined,
    VideoCameraOutlined,
} from '@ant-design/icons';
import { Button, Layout, Menu, theme } from 'antd';
import useBreakpoint from 'antd/es/grid/hooks/useBreakpoint';

const { Header, Content, Sider } = Layout;

const items = [UserOutlined, VideoCameraOutlined, UploadOutlined, UserOutlined, UserOutlined].map((icon, index) => ({
    key: String(index + 1),
    icon: React.createElement(icon),
    label: `nav ${index + 1}`,
}));

const App: React.FC = () => {
    const screens = useBreakpoint();
    const [collapsed, setCollapsed] = useState(false);

    const {
        token: { colorBgContainer, borderRadiusLG },
    } = theme.useToken();

    const handleMenuOrContent = () => {
        if (!screens.lg) {
            setCollapsed(!screens.lg);
        }
    };

    useEffect(() => {
        const layoutMarginLeft = screens.lg ? (collapsed ? 67 : 200) : 0;
        const headerMarginLeft = screens.lg ? 0 : collapsed ? 0 : 200;
        setLayoutStyle({ marginLeft: layoutMarginLeft });
        setHeaderStyle({ padding: 0, background: colorBgContainer, marginLeft: headerMarginLeft });
    }, [screens, collapsed, colorBgContainer]);

    const [layoutStyle, setLayoutStyle] = useState<{ [key: string]: string | number }>({});
    const [headerStyle, setHeaderStyle] = useState<{ [key: string]: string | number }>({});

    return (
        <Layout style={{ height: '100vh' }}>
            <Sider
                breakpoint="lg"
                collapsedWidth={screens.lg ? '' : '0'}
                onBreakpoint={(broken: boolean) => {
                    setCollapsed(broken);
                }}
                trigger={null}
                collapsible
                collapsed={collapsed}
                style={{ overflow: 'auto', height: '100vh', position: 'fixed', left: 0, top: 0, bottom: 0 }}
            >
                <div className="demo-logo-vertical" />
                <Menu
                    theme="dark"
                    mode="inline"
                    defaultSelectedKeys={['4']}
                    items={items}
                    onClick={handleMenuOrContent}
                />
            </Sider>
            <Layout style={layoutStyle}>
                <Header style={headerStyle}>
                    <Button
                        type="text"
                        icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                        onClick={() => setCollapsed(!collapsed)}
                        style={{
                            fontSize: '16px',
                            width: 64,
                            height: 64,
                        }}
                    />
                </Header>
                <Content style={{ margin: 20, overflow: 'initial' }} onClick={handleMenuOrContent}>
                    <div
                        style={{
                            height: '100%',
                            padding: 20,
                            textAlign: 'center',
                            background: colorBgContainer,
                            borderRadius: borderRadiusLG,
                        }}
                    >
                        content
                    </div>
                </Content>
            </Layout>
        </Layout>
    );
};

export default App;