export const hanlder = async (event, context) => {
    // event seria sua request
    // context seria o trace da operação
    
    const response = {
        statusCode: 200,
        body: JSON.stringify("Hello World!"),
    };

    return response;
};
