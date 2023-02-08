class Out {
  power(type, answers, explanation){
    let sentence = "";

    if(null != explanation) {
      sentence += explanation;
    
    } else {
      Logger.log(`Found ${answers.length} answers.`);

      let customResponseRepository = new CustomResponseRepository();
      let customRes = customResponseRepository.get();

      // When there is not answer found.
      if (answers.length == 0) {
        sentence = "Couldn't find an appropriate answer.";
      }
      // If there is just one answer, we send the answer's "what" or "howto" response.
      else if (answers.length == 1) {
        if("lampWhat" === type){
          sentence += answers[0].what;
        } else if("lampHowTo" === type){
          sentence += answers[0].howto;
        }
      } 
      // If there are multiple answers, we send the answer's "officialname" response.
      else if (1 < answers.length && answers.length <= customRes.upperLimitNumber) {
        sentence += "Found "
        for(let i = 0; i < answers.length; i++){
          if(1 < answers.length) {
            if(i === answers.length-1) {
              sentence += "and " + answers[i].officialname + "." + "\n" + customRes.answerMultipleLamps;
            } else {
              sentence += answers[i].officialname + ", "
            }
          }
        }
      } 
      // If there are more answers than the max amount, we'll send the number of answers in the response instead.
      else {
        sentence += `Found ${answers.length} lamps.` + "\n" + customRes.answerNumberOfLamps;
      }
    }

    // JSON response creation
    const response = {
      "sentence" : sentence + "\n"
    };

    // Log the response
    Logger.log(`out: ${sentence}`);
    
    return ContentService.createTextOutput(JSON.stringify(response));
  }
}
