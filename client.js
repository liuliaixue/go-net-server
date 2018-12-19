
// client .js
const net = require('net')
const EventEmitter = require('events')
const fs = require('fs')

const HEAD_LENGTH = 4

class Client extends EventEmitter {
  constructor(props) {
    super()
    this.buffers = [];
    console.log(props)
    this.socket = net.createConnection(props);
    this.socket.on('connect', () => this.emit('connect'))
    this.socket.on('data', buf => { this.buffers.push(buf); this.onData() });
    this.socket.on('error', err => console.log(err))
  }

  onData() {

    const data = this.buffers;
    if (data[0].length > 4) {
      const datalen = data[0].slice(0, 4).readUInt32LE()

      if (data[0].length >= datalen + 4) {

        const data1 = data[0].slice(4, 4 + datalen)
        this.onMsg(data1)
        const resetBuf = data[0].slice(4 + datalen, data[0].length)
        data[0] = resetBuf;
        this.onData(data)

      } else {
        if (data.length > 1) {
          const firstBuf = data.shift();
          data[0] = Buffer.concat([firstBuf, data[0]]);
          this.onData(data)
        }
      }

    } else if (data.length > 1) {

      const firstBuf = data.shift();
      data[0] = Buffer.concat([firstBuf, data[0]]);
      this.onData(data)
    }
  }
  onMsg(buffer) {
    const string = buffer.toString();
    try {
      console.log('@@@@@@@', JSON.parse(string));
      this.emit('message', JSON.parse(string));
    } catch (e) {

    }
  }
  writeString(string) {
    let length = Buffer.alloc(4);
    length.writeUInt32LE(string.length);
    this.socket.write(length)
    this.socket.write(string)
  }
  writeJson(json) {
    const data = JSON.stringify(json)
    let length = Buffer.alloc(4);
    length.writeUInt32LE(data.length);
    this.socket.write(length)
    this.socket.write(data)
  }
}


module.exports = Client



const executor = async () => {
  const client = new Client({ port: '8080' })
  setTimeout(() => {
    const str = 'hello from node client'
    const length = Buffer.alloc(4)
    console.log(str.length)
    length.writeUInt32LE(str.length)
    console.log(length)
    console.log(Buffer.from(str))
    client.socket.write(length)
    // client.socket.write(Buffer.from(str))
  }, 100)

  // setTimeout(() => {
  //   const data = fs.readFileSync('./pm2-out-45.log')
  //   const length = Buffer.alloc(4)
  //   length.writeUInt32LE(data.length)

  //   console.log(length)
  //   client.socket.write(length)

  // }, 200)


  // setTimeout(() => {
  //   const length = Buffer.alloc(4)
  //   length.writeUInt32LE(1 + (1 * 1 << 8) + (1 * 1 << 16) + (1 * 1 << 24))

  //   console.log(length, 1 + (1 * 1 << 8) + (1 * 1 << 16) + (1 * 1 << 24))
  //   client.socket.write(length)

  // }, 200)
}

executor()