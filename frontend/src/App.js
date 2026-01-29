import React, { useState } from 'react';
import { getSettings, getStateInstance, sendMessage, sendFileByUrl } from './services/api';

function App() {
  const [idInstance, setIdInstance] = useState('');
  const [apiTokenInstance, setApiTokenInstance] = useState('');
  const [response, setResponse] = useState('');
  const [loading, setLoading] = useState(false);

  const [chatId, setChatId] = useState('');
  const [message, setMessage] = useState('');

  const [fileChatId, setFileChatId] = useState('');
  const [fileUrl, setFileUrl] = useState('');
  const [fileName, setFileName] = useState('');

  const handleGetSettings = async () => {
    if (!idInstance || !apiTokenInstance) {
      setResponse('Ошибка: Заполните idInstance и apiTokenInstance');
      return;
    }

    setLoading(true);
    try {
      const result = await getSettings(idInstance, apiTokenInstance);
      setResponse(JSON.stringify(result, null, 2));
    } catch (error) {
      setResponse(`Ошибка: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  const handleGetStateInstance = async () => {
    if (!idInstance || !apiTokenInstance) {
      setResponse('Ошибка: Заполните idInstance и apiTokenInstance');
      return;
    }

    setLoading(true);
    try {
      const result = await getStateInstance(idInstance, apiTokenInstance);
      setResponse(JSON.stringify(result, null, 2));
    } catch (error) {
      setResponse(`Ошибка: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  const handleSendMessage = async () => {
    if (!idInstance || !apiTokenInstance) {
      setResponse('Ошибка: Заполните idInstance и apiTokenInstance');
      return;
    }

    if (!chatId || !message) {
      setResponse('Ошибка: Заполните chatId и message');
      return;
    }

    setLoading(true);
    try {
      const result = await sendMessage({
        idInstance,
        apiTokenInstance,
        chatId,
        message,
      });
      setResponse(JSON.stringify(result, null, 2));
    } catch (error) {
      setResponse(`Ошибка: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  const handleSendFileByUrl = async () => {
    if (!idInstance || !apiTokenInstance) {
      setResponse('Ошибка: Заполните idInstance и apiTokenInstance');
      return;
    }

    if (!fileChatId || !fileUrl) {
      setResponse('Ошибка: Заполните chatId и fileUrl');
      return;
    }

    setLoading(true);
    try {
      const result = await sendFileByUrl({
        idInstance,
        apiTokenInstance,
        chatId: fileChatId,
        fileUrl,
        fileName: fileName || undefined,
      });
      setResponse(JSON.stringify(result, null, 2));
    } catch (error) {
      setResponse(`Ошибка: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 p-2 md:p-4">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-2xl font-bold text-gray-800 mb-4 text-center">GreenAPI Клиент</h1>

        <div className="flex flex-col lg:flex-row gap-4">
          <div className="flex-1 space-y-3">
            <div className="bg-white rounded-lg shadow-md p-3">
              <div className="space-y-2">
                <input
                  id="idInstance"
                  type="text"
                  value={idInstance}
                  onChange={(e) => setIdInstance(e.target.value)}
                  placeholder="ID Instance"
                  className="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-green-500 focus:border-transparent"
                />
                <input
                  id="apiTokenInstance"
                  type="password"
                  value={apiTokenInstance}
                  onChange={(e) => setApiTokenInstance(e.target.value)}
                  placeholder="API Token Instance"
                  className="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-green-500 focus:border-transparent"
                />
              </div>
            </div>
            <div className="bg-white rounded-lg shadow-md p-3">
              <div className="mb-3 p-2 bg-gray-50 rounded">
                <h3 className="text-sm font-medium text-gray-700 mb-1.5">getSettings</h3>
                <button
                  onClick={handleGetSettings}
                  disabled={loading}
                  className="bg-green-500 hover:bg-green-600 text-white text-sm font-medium py-1 px-3 rounded transition duration-200 disabled:bg-gray-400 disabled:cursor-not-allowed"
                >
                  Вызвать getSettings
                </button>
              </div>
              <div className="mb-3 p-2 bg-gray-50 rounded">
                <h3 className="text-sm font-medium text-gray-700 mb-1.5">getStateInstance</h3>
                <button
                  onClick={handleGetStateInstance}
                  disabled={loading}
                  className="bg-green-500 hover:bg-green-600 text-white text-sm font-medium py-1 px-3 rounded transition duration-200 disabled:bg-gray-400 disabled:cursor-not-allowed"
                >
                  Вызвать getStateInstance
                </button>
              </div>
              <div className="mb-3 p-2 bg-gray-50 rounded">
                <h3 className="text-sm font-medium text-gray-700 mb-1.5">sendMessage</h3>
                <div className="space-y-2 mb-2">
                  <input
                    id="chatId"
                    type="text"
                    value={chatId}
                    onChange={(e) => setChatId(e.target.value)}
                    placeholder="Телефон в формате (79999999999)"
                    className="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-green-500 focus:border-transparent"
                  />
                  <textarea
                    id="message"
                    value={message}
                    onChange={(e) => setMessage(e.target.value)}
                    placeholder="Введите сообщение"
                    rows="2"
                    className="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-green-500 focus:border-transparent"
                  />
                </div>
                <button
                  onClick={handleSendMessage}
                  disabled={loading}
                  className="bg-green-500 hover:bg-green-600 text-white text-sm font-medium py-1 px-3 rounded transition duration-200 disabled:bg-gray-400 disabled:cursor-not-allowed"
                >
                  Отправить сообщение
                </button>
              </div>
              <div className="p-2 bg-gray-50 rounded">
                <h3 className="text-sm font-medium text-gray-700 mb-1.5">sendFileByUrl</h3>
                <div className="space-y-2 mb-2">
                  <input
                    id="fileChatId"
                    type="text"
                    value={fileChatId}
                    onChange={(e) => setFileChatId(e.target.value)}
                    placeholder="Телефон в формате (79999999999)"
                    className="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-green-500 focus:border-transparent"
                  />
                  <input
                    id="fileUrl"
                    type="text"
                    value={fileUrl}
                    onChange={(e) => setFileUrl(e.target.value)}
                    placeholder="URL файла"
                    className="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-green-500 focus:border-transparent"
                  />
                  <input
                    id="fileName"
                    type="text"
                    value={fileName}
                    onChange={(e) => setFileName(e.target.value)}
                    placeholder="Имя файла"
                    className="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-green-500 focus:border-transparent"
                  />
                </div>
                <button
                  onClick={handleSendFileByUrl}
                  disabled={loading}
                  className="bg-green-500 hover:bg-green-600 text-white text-sm font-medium py-1 px-3 rounded transition duration-200 disabled:bg-gray-400 disabled:cursor-not-allowed"
                >
                  Отправить файл
                </button>
              </div>
            </div>
          </div>
          <div className="flex-1 lg:max-h-screen lg:sticky lg:top-2">
            <div className="bg-white rounded-lg shadow-md p-3 h-full">
              <h2 className="text-base font-semibold text-gray-700 mb-2 pb-2 border-b border-gray-200">
                Ответ API
              </h2>
              <textarea
                value={response}
                readOnly
                placeholder="Ответы API будут отображаться здесь..."
                rows="35"
                className="w-full px-2 py-2 border border-gray-300 rounded font-mono text-xs bg-gray-50 resize-none focus:outline-none focus:ring-1 focus:ring-green-500 focus:border-transparent"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
