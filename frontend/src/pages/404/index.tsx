import { Button, Result } from 'antd';
import React from 'react';

const Error404: React.FC = () => (
    <Result
        status="404"
        title="哎呀，走丢啦"
        extra={
            <Button type="primary" size="large" onClick={() => console.log('404')}>
                返回首页
            </Button>
        }
    />
);

export default Error404;