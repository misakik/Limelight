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
          text += "<li>#{path['id']}</li>"
        if text.length > 0
          result.html text
        else
          result.html "No Result"
  else
    result.html "Please input keyword"
