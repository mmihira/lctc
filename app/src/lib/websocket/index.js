import API from 'lib/api/extensions';

function createSocket () {
  try {
    const socket = new WebSocket(API.WS_URL());
    return {socket};
  } catch (err) {
    return {err};
  }
}

export default createSocket;
