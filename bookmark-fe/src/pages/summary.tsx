import { PageContainer } from '@ant-design/pro-layout';
import {ProCard} from '@ant-design/pro-components'
import {Empty} from 'antd'
import {useEffect, useState} from "react";
import Loading from '@/components/loading'
import {lisMenu} from "@/services/menu/menu";
import {history} from 'umi'

export default () => {
  const [loading, setLoading] = useState(true)
  const [data, setData] = useState([])

  const handleFetchData = async () => {
    const result = await lisMenu()
    return result.data.list || []
  }

  useEffect(()=> {
    (async function(){
        const result = await handleFetchData()
        setData(result)
        setLoading(false)
    }())

  }, [])
  return (
    <PageContainer
      ghost
      title={false}
      fixedHeader
      header={{
        title: '书签汇总',
        breadcrumb: {
          items: [
            {
              path: '/',
              title: '首页',
            },
            {
              path: '/summary',
              title: '汇总',
            },
          ],
        },
      }}
    >
      {loading ? <Loading/> : data.length === 0 ? <div style={{marginTop: 40}}><Empty/></div> :

        <ProCard direction="column" ghost >

          {data.map((item, index)=>(
            <ProCard key={`${item.menuId}-${item.name}`} title={item.name} ghost gutter={16}  >
              {
                item.routes.map((tag, idex2)=> (
                  <ProCard key={tag.targetId} onClick={()=>{history.push(tag.path)}} colSpan={6} layout="center" bordered hoverable>
                    {tag.name}
                  </ProCard>
                ))
              }
            </ProCard>
          ))}
        </ProCard>
      }

    </PageContainer>
  )
}
