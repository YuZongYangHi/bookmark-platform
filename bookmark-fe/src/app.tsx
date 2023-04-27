import "./app.css"
import Footer from '@/components/Footer';
import type {MenuDataItem, Settings as LayoutSettings} from '@ant-design/pro-components';
import type { RunTimeLayoutConfig} from '@umijs/max';
import {Link, history} from 'umi'
import {
  ModalForm,
  ProFormText,
} from '@ant-design/pro-components';
import {Dropdown, } from 'antd'
import {EllipsisOutlined} from '@ant-design/icons'
import defaultSettings from '../config/defaultSettings';
import { errorConfig } from './requestErrorConfig';
import {PlusOutlined, TagsOutlined,TagOutlined, EditOutlined } from '@ant-design/icons';
import {Button, Form, MenuProps, message} from 'antd'
import React, {useState, useRef} from 'react';
import {
  createCategory, updateCategory,
  deleteCategory
} from "@/services/category/category";

import {lisMenu} from '@/services/menu/menu'
import {createItem, updateItem, deleteItem} from "@/services/item/item";

/**
 * @see  https://umijs.org/zh-CN/plugins/plugin-initial-state
 * */
export async function getInitialState(): Promise<{
  settings?: Partial<LayoutSettings>;
  loading?: boolean;
  menu: MenuDataItem[]
}> {

  const result = await lisMenu()
  const menuData: MenuDataItem[] = []
  if (result.success) {
    result.data.list.map(item => {
      let t = {
        path: item.path,
        name: item.name,
        icon: <TagsOutlined />,
        _id: item.menuId,
        children: [],
      }
      if (item.routes.length > 0) {
        item.routes.map(tag => {
          t.children.push({
            path: tag.path,
            name: tag.name,
            _id: tag.targetId,
            icon: <TagsOutlined />,
            _iframeUrl: tag.iframeURL
          })
        })
      }
      menuData.push(t)
    })
  }
  return {
    settings: defaultSettings as Partial<LayoutSettings>,
    menu: menuData
  };
}

// ProLayout 支持的api https://procomponents.ant.design/components/layout
export const layout: RunTimeLayoutConfig = ({ initialState, setInitialState }) => {
  const [categoryForm] = Form.useForm<{ name: string }>();
  const [categoryVisable, setCategoryVisable] = useState(false)
  const [categoryTitle, setCategoryTitle] = useState("新建分类")
  const [currentCategory, setCategoryObj] = useState({})
  const actionRef = useRef(null)
  const [itemForm] = Form.useForm<{ name: string }>();
  const [itemVisable, setItemVisable] = useState(false)
  const [itemTitle, setItemTitle] = useState("新建书签")
  const [currentItem, setItemObj] = useState({})
  const [currentCategoryId, setCurrentCategoryId] = useState(0)

  const handleDeleteCategory = async (id: number) => {
    const result = await deleteCategory(id)
    if (result.success) {
      message.success("删除成功")
      const l = initialState.menu.concat()
      l.map((item, index) => {
        if (item._id === id) {
          delete initialState.menu[index]
          actionRef.current?.reload()
          setCategoryObj({})
          if (history.location.pathname.split('/')[1] && parseInt(history.location.pathname.split('/')[1]) === id ) {
            history.push('/')
          }
          return
        }
      })
    }
  }

  const handleUpdateCategory = () => {
    categoryForm.setFieldValue("name", currentCategory.name)
    setCategoryVisable(true)
    setCategoryTitle("修改分类")
  }
  const items: MenuProps['items'] = [
    {
      key: 'create',
      label: '新增标签',
      onClick: () => {
        setItemTitle("新建书签")
        setItemVisable(true)
      }
    },
    {
      key: 'update',
      label: '修改分类',
      onClick: () => {handleUpdateCategory()},
    },
    {
      type: 'divider',
    },
    {
      key: 'delete',
      danger: true,
      label: '删除分类',
      onClick: async ()=> {await handleDeleteCategory(currentCategory._id)}
    },
  ];

  const handleTag = (item) => {
    const categoryId = item.path.split('/')[1]
    setItemObj(item)
    setItemTitle("修改书签")
    setCurrentCategoryId(parseInt(categoryId))
    itemForm.setFieldsValue({
      name: item.name,
      iframeUrl: item._iframeUrl,
    })
    setItemVisable(true)
  }
  return {
    footerRender: () => <Footer />,
    onPageChange: () => {

    },
    actionRef: actionRef,
    menuFooterRender:() => {return <></>},
    actionsRender: () => [],
    menuDataRender: (menuData: MenuDataItem[]) => {
      return initialState && initialState.menu || []
    },
    menuItemRender: (item, dom, {collapsed}) => {
      if (item.locale.split('.').length === 3) {
        let style = {width: 200}
        if (collapsed) {
          style = {}
        }
        return  <div
          style={{
            display: 'flex',
            alignItems: 'center',
            gap: 8,
            cursor: 'pointer'
          }}
        >
          <span style={style} onClick={()=>{history.push(item.path)}} ><TagOutlined/>&nbsp;&nbsp;{item.name}</span><span onClick={()=>{handleTag(item)}}><EditOutlined /></span>
        </div>
      }
      if (collapsed) {
        return <div style={{textAlign: 'center'}}><TagsOutlined/></div>
      }
     return <div
       style={{
         display: 'flex',
         alignItems: 'center',
         gap: 8,
       }}
     >
      <span style={{width: 200}}><TagsOutlined/>&nbsp;&nbsp;{item.name}</span> <Dropdown
       onOpenChange={()=>{setCategoryObj(item)}}
       menu={{ items }}
     >
       <EllipsisOutlined />
     </Dropdown>
     </div>
    },
    subMenuItemRender: (item, dom, {collapsed}) => {
      if (collapsed) {
        return <div style={{textAlign: 'center'}}> <TagsOutlined/></div>
      }
      return  <div
        style={{
          display: 'flex',
          alignItems: 'center',
          gap: 8,
        }}
      > <span style={{width: 200}}><TagsOutlined/>&nbsp;&nbsp;{item.name}</span> <Dropdown
        onOpenChange={()=>{setCategoryObj(item)}}
        menu={{ items }}
      >
        <EllipsisOutlined />
      </Dropdown></div>
    },
    menuExtraRender: ({ collapsed }) => {
      const Elem = !collapsed && (
        <div style={{paddingTop: 8, paddingBottom: 8}}>
          <Button onClick={()=>{
            setCategoryVisable(true)
            setCategoryTitle("新建分类")
          }} icon={<PlusOutlined />} block>
            Add Category
          </Button>

        </div>
      ) || (
        <div onClick={()=>{
          setCategoryVisable(true)
          setCategoryTitle("新建分类")
        }} style={{paddingTop: 8, cursor: 'pointer'}}><PlusOutlined/></div>
      )
      return <div>
        {Elem}
        <ModalForm<{
          name: string;
        }>
          title={categoryTitle}
          form={categoryForm}
          visible={categoryVisable}
          modalProps={{
            destroyOnClose: true,
            width:600,
            onCancel: () => {
              setCategoryVisable(false)
            }
          }}
          submitTimeout={2000}
          onFinish={async (values) => {
            let fetch
            if (categoryTitle === "修改分类") {
              fetch = updateCategory(currentCategory._id, values)
            } else {
              fetch = createCategory(values)
            }
            const result = await fetch
            if (result.success) {
              if (categoryTitle === "新建分类") {
                initialState.menu.push({
                  path: `/${result.data.list.id}`,
                  name: values.name,
                  icon: <TagsOutlined/>,
                  children: [],
                  _id: result.data.list.id
                })
              } else {
                const l = initialState.menu.concat()
                l.map((item, index) => {
                  if (item._id === currentCategory._id) {
                    initialState.menu[index].name = values.name
                    actionRef.current?.reload()
                    return
                  }
                })
              }
              actionRef.current?.reload()
              setCategoryVisable(false)
              return true
            }
            return false;
          }}
        >
          <ProFormText
            rules={[{required: true, message: '名称不能为空'}]}
            name="name"
            label="分类名称"
            placeholder="请输入分类名称"
          />
        </ModalForm>

        <ModalForm<{
          name: string;
        }>
          title={itemTitle}
          form={itemForm}
          visible={itemVisable}
          modalProps={{
            destroyOnClose: true,
            width:600,
            onCancel: () => {
              setItemVisable(false)
            }
          }}
          submitter={{
            render: (props, defaultDoms) => {
              return [
                itemTitle === "修改书签" &&
                <Button
                  danger
                  key="删除标签"
                  onClick={async () => {
                    const result = await deleteItem(currentItem._id)
                    if (result.success) {
                      message.success("删除成功")

                      const l = initialState.menu.concat()
                      l.map((item, index) => {
                        if (item._id === currentCategoryId) {
                          item.children.map((tag,  index2)=>{
                            if (currentItem._id == tag._id) {
                              delete initialState.menu[index].children[index2]
                            }
                          })
                        }
                      })
                    }
                    setItemVisable(false)
                    props.submit();

                    if (history.location.pathname.split('/')[2] && parseInt(history.location.pathname.split('/')[2]) === currentItem._id ) {
                      history.push('/')
                    }
                  }}
                >
                  删除标签
                </Button>,
                ...defaultDoms,
              ];
            },
          }}
          submitTimeout={2000}
          onFinish={async (values) => {
            let fetch
            if (itemTitle === "新建书签") {
              values.categoryId = currentCategory._id
              fetch = createItem(values)
            } else {
              fetch = updateItem(currentItem._id, values)
            }
            const result = await fetch
            if (result.success) {
              if (itemTitle === "新建书签") {
                const l = initialState.menu.concat()
                l.map((item, index) => {
                  if (item._id === currentCategory._id) {
                    initialState.menu[index].children.push({
                      path: `/${currentCategory._id}/${result.data.list.id}`,
                      name: result.data.list.name,
                      _id: result.data.list.id,
                      icon: <TagsOutlined />,
                      _iframeUrl: result.data.list.iframeURL
                    })
                    history.push( `/${currentCategory._id}/${result.data.list.id}`)
                    return
                  }
                })
              }else {
                const l = initialState.menu.concat()
                l.map((item, index) => {
                  if (item._id === currentCategoryId) {
                    item.children.map((tag,  index2)=>{
                      if (currentItem._id == tag._id) {
                        initialState.menu[index].children[index2].name = result.data.list.name
                        let mark = false
                        if (initialState.menu[index].children[index2]._iframeUrl != result.data.list.iframeURL){
                          mark = true
                        }
                        initialState.menu[index].children[index2]._iframeUrl = result.data.list.iframeURL
                        if (mark) {
                          window.location.reload()
                        }
                      }
                    })
                  }
                })
              }
              setItemVisable(false)
            }
              return true
          }}
        >
          <ProFormText
            rules={[{required: true, message: '名称不能为空'}]}
            name="name"
            label="书签名称"
            placeholder="请输入书签名称"
          />
          <ProFormText
            rules={[{required: true, message: '关联链接不能为空'}]}
            name="iframeUrl"
            label="关联链接"
            placeholder="请输入关联链接"
          />
        </ModalForm>

      </div>
    },
    links: [],
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    // 增加一个 loading 的状态
    childrenRender: (children) => {
      // if (initialState?.loading) return <PageLoading />;
      return (
        <>
          {children}
        </>
      );
    },
    ...initialState?.settings,
  };
};

/**
 * @name request 配置，可以配置错误处理
 * 它基于 axios 和 ahooks 的 useRequest 提供了一套统一的网络请求和错误处理方案。
 * @doc https://umijs.org/docs/max/request#配置
 */
export const request = {
  ...errorConfig,
};
function setMenuStyle(copy: any) {
    throw new Error("Function not implemented.");
}

