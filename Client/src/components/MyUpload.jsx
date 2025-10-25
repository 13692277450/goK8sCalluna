import React, { useState } from 'react';
import { BulbFilled, LoadingOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, Upload } from 'antd';

const MyUpload = ({ onFileContentChange }) => {
  const [loading, setLoading] = useState(false);

  const beforeUpload = file => {
    const isYaml = file.type === 'application/x-yaml' || file.name.endsWith('.yaml') || file.name.endsWith('.yml');
    if (!isYaml) {
      message.error('You can only upload YAML file!');
      return false;
    }
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isLt2M) {
      message.error('YAML must smaller than 2MB!');
      return false;
    }
    return true;
  };

  const handleChange = info => {
    if (info.file.originFileObj) {
      setLoading(true);
      const reader = new FileReader();
      reader.onload = (e) => {
        setLoading(false);
        onFileContentChange(e.target.result);
      };
      reader.onerror = () => {
        setLoading(false);
        message.error('File reading failed');
      };
      reader.readAsText(info.file.originFileObj);
    }
  };

  const uploadButton = (
    <Button
      icon={loading ? <LoadingOutlined /> : <PlusOutlined />}
      style={{
        display: 'flex',
        backgroundColor: "yellowgreen",
        alignItems: 'center',
        justifyContent: 'center',
        height: '32px', // 与Modal按钮高度一致
        padding: '0 16px', // 与Modal按钮内边距一致
      }}
    >
      Upload Yaml File
    </Button>
  );

  return (
    <Upload
      name="yamlFile"
      showUploadList={false}
      beforeUpload={beforeUpload}
      onChange={handleChange}
      accept=".yaml,.yml"
      customRequest={({ onSuccess }) => {
        // 立即执行上传成功回调，因为我们直接处理文件
        setTimeout(() => onSuccess("ok"), 0);
      }}
    >
      {uploadButton}
    </Upload>
  );
}

export default MyUpload;