import * as messages from './messages.json';

const URL = 'http://54.226.90.220:3000'; // Replace with the EC2 instance public IP

async function main() {
  try {
    let offset = 0;
    const limit = 20;

    while (offset < messages.length) {
      const chunk = messages.slice(offset, offset + limit);
      const promises = chunk.map((message) => {
        return fetch(`${URL}`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            message,
          }),
        });
      });

      await Promise.all(promises);
      offset += limit;
    }
  } catch (err) {
    console.error(err);
  }
}

main();
