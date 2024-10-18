function jsonResponse(json, status = 200) {
  return new Response(JSON.stringify(json), {
      status: status,
      headers: {
          'Content-Type': 'application/json'
      }
  });
}

export async function GET({ url }) {
  const isPrivate = url.searchParams.get('is_private');
  let apiUrl = '';
  if (isPrivate == 'true' || isPrivate == 'false') {
    apiUrl = `http://api:8000/notes?is_private=${isPrivate}`;
  } else {
    apiUrl = `http://api:8000/notes`;
  }
  
  try {
    const response = await fetch(apiUrl);
    if (response.ok) {
      const data = await response.json();
      return jsonResponse(data);
    
    } else {
      return jsonResponse({status: response.status, error: 'Failed to fetch data!'});
    }

  } catch (error) {
    return jsonResponse({status: 500, error: 'Internal server error'});
  }
}

export async function POST({ url }) {
  const title = url.searchParams.get('title');
  const noteBody = url.searchParams.get('note_body');
  let isPrivate = url.searchParams.get('is_private');

  if (isPrivate != 'true' && isPrivate != "false") {
    isPrivate = false;
  }

  const apiUrl = `http://api:8000/notes`;

  try {
    const response = await fetch(apiUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        title: title,
        note_body: noteBody,
        is_private: isPrivate
      })
    });

    const data = await response.json();
    if (response.ok) {
      return jsonResponse([data]);

    } else if (response.status == 400){
      return jsonResponse({status: response.status, error: data.detail.msg});

    } else {
      return jsonResponse({status: response.status, error: 'Failed to fetch data!'});
    }
  } catch (error) {
    return jsonResponse({status: 500, error: 'Internal server error'});
  }
}

export async function PUT({ url }) {
  const noteID = url.searchParams.get('note_id');
  const title = url.searchParams.get('title');
  const noteBody = url.searchParams.get('note_body');
  let isPrivate = url.searchParams.get('is_private');

  if (isPrivate != 'true' && isPrivate != "false") {
    isPrivate = false;
  }

  if (isNaN(noteID)) {
    return jsonResponse({status: 400, error: 'Attribute `id` must be a number!'});
  }

  const apiUrl = `http://api:8000/notes/${noteID}`;

  try {
    const response = await fetch(apiUrl, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        title: title,
        note_body: noteBody,
        is_private: isPrivate
      })
    });

    const data = await response.json();
    if (response.ok) {
      return jsonResponse([data]);

    } else if (response.status == 400){
      return jsonResponse({status: response.status, error: data.detail.msg});

    } else if (response.status == 404){
      return jsonResponse({status: response.status, error: data.detail.msg});

    } else {
      return jsonResponse({status: response.status, error: 'Failed to fetch data!'});
    }
  } catch (error) {
    return jsonResponse({status: 500, error: 'Internal server error'});
  }
}

export async function DELETE({ url }) {
  const noteID = url.searchParams.get('note_id');
  
  if (isNaN(noteID)) {
    return jsonResponse({status: 400, error: 'Attribute `id` must be a number!'});
  }

  const apiUrl = `http://api:8000/notes/${noteID}`;

  try {
    const response = await fetch(apiUrl, {
      method: 'DELETE'
    });

    const data = await response.json();
    if (response.ok) {
      return jsonResponse(data);

    } else if (response.status == 404){
      return jsonResponse({status: response.status, error: data.detail.msg});

    } else {
      return jsonResponse({status: response.status, error: 'Failed to fetch data!'});
    }
  } catch (error) {
    return jsonResponse({status: 500, error: 'Internal server error'});
  }
}