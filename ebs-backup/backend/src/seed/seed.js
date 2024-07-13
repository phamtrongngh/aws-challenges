const URL = 'http://54.174.185.15:3000';

async function postMessage(message) {
  try {
    const res = await fetch(URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ message }),
    });

    if (!res.ok) {
      throw new Error('Failed to post message');
    }
    console.log(`Posted message: ${message}`);
  } catch (err) {
    console.error(err);
  }
}

async function main() {
  try {
    const res = await fetch(
      'https://randomwordgenerator.com/json/sentences.json',
    );

    const body = await res.json();
    const { data } = body;

    let offset = 0;
    let limit = 20;

    while (offset < data.length) {
      const messages = data.slice(offset, offset + limit);
      const promises = messages.map((message) => postMessage(message));
      await Promise.all(promises);
      offset += limit;
    }
  } catch (err) {
    console.error(err);
  }
}

main();
