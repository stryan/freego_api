import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import Board from '../components/board'
import styles from '../styles/Home.module.css'

const Home: NextPage = () => (
  <>
    <Head>
      <title>Free Go Game</title>
      <link rel="icon" href="/favicon.ico" />
    </Head>
    <main>
      <Board />
    </main>
  </>
)

export default Home