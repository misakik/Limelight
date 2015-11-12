search = $('input.search')
result = $('ul.result')

search.on 'keyup', ->
  text = ""
  keyword = $(@).val()
  if keyword != ""
    $.ajax
      url: 'http://localhost:8000/search/' + keyword
      dataType: 'json'
      success: (data) ->
        text = ""
        for path in data['hits']
          text += "<li><span class=\"score\">#{path['score']}</span> : #{path['id']}</li>"
        if text.length > 0
          total_hits = data['total_hits']
          took = data['took']
          text = "<div class=\"hits\">#{total_hits} hits, #{took} μsec</div>" + text
          result.html text
        else
          result.html "No Result"
  else
    result.html "Please input keyword"
