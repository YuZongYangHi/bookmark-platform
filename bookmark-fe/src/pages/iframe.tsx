import React, {useEffect, useState} from 'react';
import { GridContent } from '@ant-design/pro-layout';
import {useParams} from 'umi'
import {getItem} from '@/services/item/item'

function IframeComponent() {
  const [iFrameHeight, setiFrameHeight] = useState('0px');
  const [iframeURL, setIframeURL] = useState("")

  const params = useParams()

  const fetch = async () => {
    const result = await getItem(params.itemId)
    if (result.success) {
      return result.data.list.iframeURL
    }
    return ""
  }

  useEffect(()=> {
    (async function(){
      const iframe = await fetch()
      setIframeURL(iframe)
    }())
  }, [params.itemId])

  return (
    <GridContent>
      <div id="container" style={{width: "100%", height: "100%", overflowX: "scroll"}}>
        <iframe
          frameBorder="0"
          sandbox="allow-same-origin allow-scripts allow-popups allow-forms"
          scrolling="auto"
          style={{ width: '100%', height: iFrameHeight, overflow: 'visible' }}
          onLoad={() => {
            console.log(1);
            const h = document.documentElement.clientHeight - 20;
            setiFrameHeight(`${h}px`);
          }}
          src={iframeURL}
        />
      </div>
    </GridContent>
  );
};

export default IframeComponent
