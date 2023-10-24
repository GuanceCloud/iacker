import { Cards  } from 'nextra/components'
import './index.module.css'

export const Integrations = (props) => {
    return <Cards {...props} />
}

export const Item = ({children, ...props}) => {
    const cardProps = {...props, ...{
        title: props.title || '未知',
        href: props.href || '',
        icon: props.icon || '',
        image: true,
        arrow: true,
        target: '_blank',
    }}
    return <Cards.Card {...cardProps}>
        <div className="bg-white">{children}</div>
    </Cards.Card>
}
