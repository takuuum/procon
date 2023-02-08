class Out {
  power(type, answers, explanation){
    let sentence = "";

    if(null != explanation) {
      sentence += explanation;
    
    } else {
      Logger.log(`Found ${answers.length} answers.`);

      if (1 < answers.length) {
        sentence += "Found "
      }
      for(let i = 0; i < answers.length; i++){
        if(1 < answers.length) {
          if(i === answers.length-1) {
            sentence += "and " + answers[i].officialname + "."
          } else {
            sentence += answers[i].officialname + ", "
          }
        }
        // If there is just one answer, we send the answer's "what" or "howto" response.
        else {
          if("lampWhat" === type){
              sentence += answers[i].what;
          } else if("lampHowTo" === type){
              sentence += answers[i].howto;
          }
        }
      }

      // When there is not answer found.
      if(answers.length === 0 ){
        sentence = "Couldn't find an appropriate answer."
      }
    }

    // JSON response creation
    const response = {
      "sentence" : sentence
    };

    // Log the response
    Logger.log(`out: ${sentence}`);
    
    return ContentService.createTextOutput(JSON.stringify(response));
  }
}
